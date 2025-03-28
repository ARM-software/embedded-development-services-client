package codegen

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"path/filepath"
	"strings"

	_ "github.com/ARM-software/embedded-development-services-client/client"
	"github.com/ARM-software/golang-utils/utils/commonerrors"
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/iancoleman/strcase"
	"golang.org/x/tools/go/packages"
)

type APIServiceFollowFunc struct {
	Name           string
	FirstParamName string
	FirstParamType string
	ReturnList     string
	Src            string
}

type Follower struct {
	APIService    string
	APIRequest    string
	APIFollowFunc APIServiceFollowFunc
}

type Followers = []Follower

type FollowersParams = struct {
	Followers
}

const followFuncComment = `// Follows a collection link. This function is based on the ` + "`%s`" + ` function, with one key difference:
// instead of using a fixed endpoint path, ` + "`localVarPath`" + ` is defined as the base URL + the ` + "`link`" + ` provided in the arguments.`

func AddLinkFollowersToParams(d *Data) error {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) { return getLinkFollowersParams(swagger) }, "linkfollowers.go")
}

func getLinkFollowersParams(swagger *openapi3.T) (params FollowersParams, err error) {
	ids, idsErr := getOperationIDs(swagger)
	if idsErr != nil {
		err = idsErr
		return
	}

	funcNameMap := getFuncNameReplacementMap(ids)

	followers, followersErr := getLinkFollowers(funcNameMap)
	if followersErr != nil {
		err = followersErr
		return
	}

	params = FollowersParams{
		followers,
	}

	return
}

func generateExecuteFnName(in string) string {
	// 'ToCamel()' mimics openapi-generator's 'camelize(sanitizeName(operationId))' logic for transforming OperationID.
	// The "Execute" suffix is hardcoded in the template,
	// see https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator/src/main/resources/go/api.mustache
	return strcase.ToCamel(in) + "Execute"
}

func generateFollowFuncName(in string) string {
	return "Follow" + strings.TrimSuffix(strings.TrimPrefix(strcase.ToCamel(in), "List"), "Execute") + "Link"
}

func getFuncNameReplacementMap(operationIDs mapset.Set[string]) map[string]string {
	executeFuncList := Map(operationIDs.ToSlice(), generateExecuteFnName)
	funcNameMap := make(map[string]string)
	for _, fnName := range executeFuncList {
		funcNameMap[fnName] = generateFollowFuncName(fnName)
	}

	return funcNameMap
}

func getOperationIDs(swagger *openapi3.T) (opIDs mapset.Set[string], err error) {
	opIDs = mapset.NewSet[string]()

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

					opIDs.Add(pathInfo.Get.OperationID)
				}

				if isMessagesCollection := strings.HasSuffix(endpoint, "messages"); isMessagesCollection {
					if strings.Contains(collectionRef, notificationFeedRef) || strings.Contains(collectionRef, messageItemRef) {
						if shouldIgnoreItem(collectionRef) {
							continue
						}

						opIDs.Add(pathInfo.Get.OperationID)
					}
				}
			}
		}
	}

	return
}

func renderSrc(fn *ast.FuncDecl, fset *token.FileSet) (src string, err error) {
	var buf bytes.Buffer
	renderErr := printer.Fprint(&buf, fset, fn)
	if renderErr != nil {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not render source code for func '%s()'", fn.Name.Name)
		return
	}

	src = buf.String()
	return
}

func renderFieldListSrc(fields *ast.FieldList, fset *token.FileSet) (src string, err error) {
	if fields == nil || len(fields.List) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not render source code for FieldList, no fields exists")
		return
	}

	// Return just one type with no paranthesis or name
	var buf bytes.Buffer
	if len(fields.List) == 1 && len(fields.List[0].Names) == 0 {
		printer.Fprint(&buf, fset, fields.List[0].Type)
		src = buf.String()
		return
	}

	// Process multiple return types, named and unnamed
	buf.WriteByte('(')
	for i, field := range fields.List {
		for j, name := range field.Names {
			buf.WriteString(name.Name)
			if j < len(field.Names)-1 {
				buf.WriteByte(',')
			} else {
				buf.WriteByte(' ')
			}
		}

		printer.Fprint(&buf, fset, field.Type)

		if i < len(fields.List)-1 {
			buf.WriteString(", ")
		}
	}
	buf.WriteByte(')')

	src = buf.String()
	return
}

func modifyFirstParam(params *ast.FieldList) (err error) {
	if params == nil || len(params.List) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not modify first parameter, no fields exists")
		return
	}

	param := params.List[0]

	if _, ok := param.Type.(*ast.StarExpr); !ok {
		param.Type = &ast.StarExpr{
			Star: token.NoPos,
			X:    param.Type,
		}
	}

	return
}

func getFirstParam(params *ast.FieldList, fset *token.FileSet) (paramName string, paramType string, err error) {
	if params == nil || len(params.List) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not render source code for first parameter, no fields exists")
		return
	}

	param := params.List[0]

	if len(param.Names) > 1 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not get first parameter, expected only one name")
		return
	}

	if len(param.Names) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not get first parameter, expected a named parameter")
		return
	}

	paramName = param.Names[0].Name
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, param.Type)
	paramType = buf.String()

	return
}

func getReceiverType(fn *ast.FuncDecl) (reciever string, err error) {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "No receiver defined for func `%s()`", fn.Name.Name)
		return
	}

	expr := fn.Recv.List[0].Type

	switch t := expr.(type) {
	case *ast.Ident:
		reciever = t.Name
		return
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			reciever = ident.Name
			return
		}
	}

	err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not get receiver for func `%s()`", fn.Name.Name)
	return
}

func modifyLocalVarPath(fn *ast.FuncDecl) (err error) {
	for _, stmt := range fn.Body.List {
		assign, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}

		if assign.Tok == token.DEFINE && len(assign.Lhs) > 0 {
			if ident, ok := assign.Lhs[0].(*ast.Ident); ok && ident.Name == "localVarPath" {
				if expr, ok := assign.Rhs[0].(*ast.BinaryExpr); ok && len(assign.Rhs) > 0 {
					expr.Y = ast.NewIdent("link")
				}

				return
			}
		}
	}

	err = commonerrors.Newf(commonerrors.ErrUnexpected, "Could not find `localVarPath` definition")
	return
}

func addFuncDocs(fn *ast.FuncDecl) {
	fn.Doc = &ast.CommentGroup{
		List: []*ast.Comment{
			{
				Slash: fn.Pos() - 1,
				Text:  fmt.Sprintf(followFuncComment, fn.Name.Name),
			},
		},
	}
}

func addNewFuncParam(fn *ast.FuncDecl, argName string, argType string) {
	fn.Type.Params.List = append(fn.Type.Params.List, &ast.Field{
		Names: []*ast.Ident{ast.NewIdent(argName)},
		Type:  ast.NewIdent(argType),
	})
}

func getLinkFollowers(FuncNameMap map[string]string) (followers Followers, err error) {
	// 'parser.ParseDir()' returns deprecated 'ast.Package', so we'll use 'packages.Load()' till it gets updated to return 'packages.Package'
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedSyntax | packages.NeedCompiledGoFiles,
		Tests: false,
	}
	pkgs, loadErr := packages.Load(cfg, "github.com/ARM-software/embedded-development-services-client/client")
	if loadErr != nil {
		err = commonerrors.Newf(commonerrors.ErrInvalidDestination, "Could not load packages from directory '%s'", cfg.Dir)
		return
	}

	for _, pkg := range pkgs {
		for i, file := range pkg.Syntax {
			// Filter filenames that match "api_*.go"
			filename := pkg.CompiledGoFiles[i]
			if !strings.HasPrefix(filepath.Base(filename), "api_") {
				continue
			}

			for _, decl := range file.Decls {
				if fn, ok := decl.(*ast.FuncDecl); ok {
					if followFuncName, ok := FuncNameMap[fn.Name.Name]; ok {
						// Process APIService.FollowLink function
						addFuncDocs(fn)
						fn.Name.Name = followFuncName
						addNewFuncParam(fn, "link", "string")
						modifyParamErr := modifyFirstParam(fn.Type.Params)
						if modifyParamErr != nil {
							err = modifyParamErr
						}
						modifyLineErr := modifyLocalVarPath(fn)
						if modifyLineErr != nil {
							err = modifyLineErr
						}
						APIServiceFollowFuncSrc, renderErr := renderSrc(fn, pkg.Fset)
						if renderErr != nil {
							err = renderErr
						}
						APIServiceRecieverType, renderErr := getReceiverType(fn)
						if renderErr != nil {
							err = renderErr
						}

						// Render out parts of the Request.FollowLink funciton
						paramName, paramType, renderErr := getFirstParam(fn.Type.Params, pkg.Fset)
						if renderErr != nil {
							err = renderErr
						}
						requestReturns, renderErr := renderFieldListSrc(fn.Type.Results, pkg.Fset)
						if renderErr != nil {
							err = renderErr
						}

						followers = append(followers, Follower{
							APIService: APIServiceRecieverType,
							APIRequest: strings.TrimPrefix(paramType, "*"),
							APIFollowFunc: APIServiceFollowFunc{
								Name:           fn.Name.Name,
								FirstParamName: paramName,
								FirstParamType: paramType,
								ReturnList:     requestReturns,
								Src:            APIServiceFollowFuncSrc,
							},
						})
					}
				}
			}
		}
	}

	return
}
