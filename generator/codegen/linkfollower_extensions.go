package codegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"path/filepath"
	"regexp"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

type APIServiceFollowFunc struct {
	Name       string
	FirstParam string
	ReturnList string
	Src        string
}

type Follower struct {
	APIService    string
	APIFollowFunc APIServiceFollowFunc
}

type Followers = []Follower

type FollowersParams = struct {
	Followers
}

const followFuncComment = `// Follows a collection link. This function is based on the ` + "`%s`" + ` function, with one key difference:
instead of using a fixed endpoint path, ` + "`localVarPath`" + ` is constructed by combining the base URL with the ` + "`link`" + ` provided in the arguments.`

func Map[T, V comparable](input mapset.Set[T], mapfn func(T) V) mapset.Set[V] {
	output := mapset.NewSet[V]()
	for t := range input.Iter() {
		output.Add(mapfn(t))
	}
	return output
}

func generateExecuteFnName(in string) (out string) {
	// 'ToCamel()' mimics openapi-generator's 'camelize(sanitizeName(operationId))' logic for transforming OperationID.
	// The "Execute" suffix is hardcoded in the template,
	// see https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator/src/main/resources/go/api.mustache
	return strcase.ToCamel(in) + "Execute"
}

func generateFollowFnName(in string) (out string) {
	return "Follow" + strings.TrimSuffix(strings.TrimPrefix(strcase.ToCamel(in), "List"), "Execute") + "Link"
}

func GenerateFollowLinks(d *Data) error {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) {
		ids, _ := getOperationIDs(swagger)
		executeFunctions := Map(ids, generateExecuteFnName)
		executeFunctionsList := executeFunctions.ToSlice()
		replacementFuncMap := make(map[string]string)
		for _, fnName := range executeFunctionsList {
			replacementFuncMap[fnName] = generateFollowFnName(fnName)
		}

		// FollowFunctionsList := FollowFunctions.ToSlice()
		for k, v := range replacementFuncMap {
			fmt.Printf("Execute: %s(), Replacement: %s()\n", k, v)
		}
		return getFuncs(replacementFuncMap)

		// for _, f := range funcs {
		// 	fmt.Println(f.Name.Name)
		// }
	}, "linkfollowers.go")
}

func ReplaceLocalVarPathLine(source string, originalFuncName string) string {
	re := regexp.MustCompile(`(?m)localVarPath\s*:=\s*localBasePath\s*\+\s*".*"$`)

	replacement := `
	// NOTE:
	// These lines are the only differences from the ` + "`" + originalFuncName + "`" + ` function,
	// we specify the destination of the call by appending the link that is to be followed
	localVarPath := localBasePath + link
	// **************************************************************************************************************
	`

	return re.ReplaceAllString(source, replacement)
}

func getOperationIDs(swagger *openapi3.T) (opids mapset.Set[string], err error) {
	opids = mapset.NewSet[string]()

	for endpoint, pathInfo := range swagger.Paths {
		var isCollection bool

		// if no get field then not a collection
		if pathInfo.Get != nil {
			response200Val, subErr := getResponseFromPathInfo(pathInfo, endpoint, 200)
			if subErr != nil {
				err = subErr
				return
			}

			var isRedacted bool
			if c, ok := pathInfo.Get.ExtensionProps.Extensions[redactFlag].(json.RawMessage); ok {
				err = json.Unmarshal(c, &isRedacted)
				if err != nil {
					return
				}
			}
			if isRedacted {
				continue
			}

			// if get not "application/json" then it is "application/octet-stream" and collections don't have that
			if appJSON := response200Val.Content.Get(jsonMIME); appJSON != nil {
				schemaVal, collectionRef, subErr := getCollectionSchema(swagger, appJSON, endpoint)
				if subErr != nil {
					err = subErr
					return
				}

				if c, ok := schemaVal.ExtensionProps.Extensions[collectionFlag].(json.RawMessage); ok {
					err = json.Unmarshal(c, &isCollection)
					if err != nil {
						return
					}

					if !isCollection || shouldIgnoreCollection(collectionRef) {
						continue
					}

					itemRef, subErr := extractItemRef(schemaVal, collectionRef)
					if subErr != nil {
						err = subErr
						return
					}

					if shouldIgnoreItem(itemRef) {
						continue
					}

					opids.Add(pathInfo.Get.OperationID)
				}

				if isMessagesCollection := strings.HasSuffix(endpoint, "messages"); isMessagesCollection {
					if strings.Contains(collectionRef, notificationFeedRef) || strings.Contains(collectionRef, messageItemRef) {
						if shouldIgnoreItem(collectionRef) {
							continue
						}

						opids.Add(pathInfo.Get.OperationID)
					}
				}
			}
		}
	}

	return
}

func generateFuncDeclScr(fn *ast.FuncDecl, fset *token.FileSet) (string, error) {
	var buf bytes.Buffer
	err := printer.Fprint(&buf, fset, fn)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}

func getReceiverName(fn *ast.FuncDecl) string {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		return ""
	}

	expr := fn.Recv.List[0].Type

	switch t := expr.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			return ident.Name
		}
	}

	return ""
}

func modifyLocalVarPath(fn *ast.FuncDecl) {
	// Walk the function's body statements.
	for i, stmt := range fn.Body.List {
		assign, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}

		// Look for a short variable declaration (":=") where the first LHS identifier is "localVarPath".
		if assign.Tok == token.DEFINE && len(assign.Lhs) > 0 {
			if ident, ok := assign.Lhs[0].(*ast.Ident); ok && ident.Name == "localVarPath" {
				// Create the new right-hand side: localBasePath + link.
				newRhs := &ast.BinaryExpr{
					X:  ast.NewIdent("localBasePath"),
					Op: token.ADD,
					Y:  ast.NewIdent("link"),
				}

				// Build the new assignment statement.
				newAssign := &ast.AssignStmt{
					Lhs: assign.Lhs,
					Tok: assign.Tok,
					Rhs: []ast.Expr{newRhs},
				}

				// Replace the old assignment with the new one.
				fn.Body.List[i] = newAssign
				break
			}
		}
	}
}

func getFuncs(replacementFuncMap map[string]string) (params FollowersParams, err error) {
	// 'parser.ParseDir()' returns deprecated 'ast.Package', so we'll use 'packages.Load()' till it gets updated to return 'packages.Package'
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedSyntax | packages.NeedCompiledGoFiles,
		Tests: false,
		Dir:   "../client",
	}
	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return
	}

	var followers Followers
	for _, pkg := range pkgs {
		for i, file := range pkg.Syntax {
			// Filter filenames that match "api_*.go"
			filename := pkg.CompiledGoFiles[i]
			if !strings.HasPrefix(filepath.Base(filename), "api_") {
				continue
			}

			println(filename)
			var APIServiceExecuteFunc *ast.FuncDecl
			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {

					if _, ok := replacementFuncMap[fn.Name.Name]; ok {
						APIServiceExecuteFunc = fn
						println("\tFOUND: ", fn.Name.Name, ", ReplacementName: ", replacementFuncMap[fn.Name.Name])

						originalAPIName := APIServiceExecuteFunc.Name.Name
						comments := APIServiceExecuteFunc.Doc.List
						if len(comments) > 0 {
							comments[0].Text = fmt.Sprintf(followFuncComment, originalAPIName)
						}
						newFuncName := replacementFuncMap[APIServiceExecuteFunc.Name.Name]
						APIServiceExecuteFunc.Name.Name = newFuncName
						APIServiceExecuteFunc.Type.Params.List = append(APIServiceExecuteFunc.Type.Params.List, &ast.Field{
							Names: []*ast.Ident{ast.NewIdent("link")},
							Type:  ast.NewIdent("string"),
						})
						// reqExecuteFuncSrc, _ := renderFuncDeclToString(reqExecuteFunc)
						modifyLocalVarPath(APIServiceExecuteFunc)
						APIServiceExecuteFuncSrc, _ := generateFuncDeclScr(APIServiceExecuteFunc, pkg.Fset)
						// APIServiceExecuteFuncSrc = ReplaceLocalVarPathLine(APIServiceExecuteFuncSrc, originalAPIName)

						var buf bytes.Buffer
						fset := token.NewFileSet()
						err = printer.Fprint(&buf, fset, APIServiceExecuteFunc.Type)
						if err != nil {
							fmt.Println(err.Error())
							return
						}
						funcTypeSrc := buf.String()
						re := regexp.MustCompile(`^func(\((?P<param>[^,]*).*\)) (?P<returns>\(.*\))$`)
						// if re.MatchString(funcTypeSrc) {
						matches := re.FindStringSubmatch(funcTypeSrc)
						param := matches[re.SubexpIndex("param")]
						returns := matches[re.SubexpIndex("returns")]
						// }

						followers = append(followers, Follower{
							APIService: getReceiverName(APIServiceExecuteFunc),
							APIFollowFunc: APIServiceFollowFunc{
								Name:       newFuncName,
								FirstParam: "(" + param + ")",
								ReturnList: returns,
								Src:        APIServiceExecuteFuncSrc,
							},
						})
					}
				}
			}
		}
	}

	params = FollowersParams{
		followers,
	}
	return
}
