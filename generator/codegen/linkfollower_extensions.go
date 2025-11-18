package codegen

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/printer"
	"go/token"
	"path/filepath"
	"strings"

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

const (
	resourcePath          = "link"
	pagingLimitField      = "limit"
	pagingLimitExistFlag  = "linkHasLimitParam"
	pagingOffsetField     = "offset"
	pagingOffsetExistFlag = "linkHasOffsetParam"
	pagingEmbedField      = "embed"
	pagingEmbedExistFlag  = "linkHasEmbedParam"
)

func AddLinkFollowersToParams(d *Data) error {
	return AddValuesToParams(d, func(swagger *openapi3.T) (interface{}, error) { return getLinkFollowersParams(swagger, d) }, "linkfollowers.go")
}

func getLinkFollowersParams(swagger *openapi3.T, d *Data) (params FollowersParams, err error) {
	ids, idsErr := getListingOperationIDs(swagger)
	if idsErr != nil {
		err = idsErr
		return
	}

	funcNameMap := make(map[string]string)
	for _, opID := range ids.ToSlice() {
		// 'strcase.ToCamel' mimics openapi-generator's 'camelize(sanitizeName(operationId))' logic for transforming OperationID.
		// The "Execute" suffix is hardcoded in the template,
		// see https://github.com/OpenAPITools/openapi-generator/blob/master/modules/openapi-generator/src/main/resources/go/api.mustache
		opID = strcase.ToCamel(opID)
		execFuncName := fmt.Sprintf("%sExecute", opID)
		followFuncName := fmt.Sprintf("Follow%sLink", strings.TrimPrefix(opID, "List"))
		funcNameMap[execFuncName] = followFuncName
	}

	followers, followersErr := getLinkFollowers(funcNameMap, d)
	if followersErr != nil {
		err = followersErr
		return
	}

	params = FollowersParams{
		followers,
	}

	return
}

func getListingOperationIDs(swagger *openapi3.T) (opIDs mapset.Set[string], err error) {
	opIDs = mapset.NewSet[string]()

	for endpoint, pathInfo := range swagger.Paths {
		// if no get field then not a collection
		if pathInfo.Get != nil {
			response200Val, subErr := getResponseFromPathInfo(pathInfo, endpoint, 200)
			if subErr != nil {
				err = subErr
				return
			}

			isRedacted, subErr := isExtensionFlagSet(pathInfo.Get.ExtensionProps, redactFlag)
			if subErr != nil {
				err = subErr
				return
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

				isCollection, subErr := isExtensionFlagSet(schemaVal.ExtensionProps, collectionFlag)
				if subErr != nil {
					err = subErr
					return
				}

				if isCollection {
					if shouldIgnoreCollection(collectionRef) {
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
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not render source code for func '%s'", fn.Name.Name)
		return
	}

	src = buf.String()
	return
}

func renderFieldListSrc(fields *ast.FieldList, fset *token.FileSet) (src string, err error) {
	if fields == nil || len(fields.List) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not render source code for FieldList, no fields exists")
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
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not modify first parameter, no fields exists")
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
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not render source code for first parameter, no fields exists")
		return
	}

	param := params.List[0]

	if len(param.Names) > 1 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not get first parameter, expected only one name")
		return
	}

	if len(param.Names) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not get first parameter, expected a named parameter")
		return
	}

	paramName = param.Names[0].Name
	var buf bytes.Buffer
	printer.Fprint(&buf, fset, param.Type)
	paramType = buf.String()

	return
}

func getReceiverType(fn *ast.FuncDecl) (receiver string, err error) {
	if fn.Recv == nil || len(fn.Recv.List) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "no receiver defined for func `%s`", fn.Name.Name)
		return
	}

	expr := fn.Recv.List[0].Type

	switch t := expr.(type) {
	case *ast.Ident:
		receiver = t.Name
		return
	case *ast.StarExpr:
		if ident, ok := t.X.(*ast.Ident); ok {
			receiver = ident.Name
			return
		}
	}

	err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not get receiver for func `%s`", fn.Name.Name)
	return
}

func createEmptyLinkCheck() *ast.IfStmt {
	// link == ""
	cond := &ast.BinaryExpr{
		X:  &ast.Ident{Name: "link"},
		Op: token.EQL,
		Y:  &ast.BasicLit{Kind: token.STRING, Value: "\"\""},
	}

	// errors.New("empty link")
	callExpr := &ast.CallExpr{
		Fun: &ast.SelectorExpr{
			X:   &ast.Ident{Name: "errors"},
			Sel: &ast.Ident{Name: "New"},
		},
		Args: []ast.Expr{
			&ast.BasicLit{Kind: token.STRING, Value: "\"empty link\""},
		},
	}

	// return nil, nil, errors.New("empty link")
	retStmt := &ast.ReturnStmt{
		Results: []ast.Expr{
			&ast.Ident{Name: "nil"},
			&ast.Ident{Name: "nil"},
			callExpr,
		},
	}

	// { return nil, nil, errors.New("empty link") }
	block := &ast.BlockStmt{
		List: []ast.Stmt{retStmt},
	}

	// if link == "" {
	//     return nil, nil, errors.New("empty link")
	// }
	ifStmt := &ast.IfStmt{
		Cond: cond,
		Body: block,
	}

	return ifStmt
}

func modifyLocalVarPath(fn *ast.FuncDecl) (err error) {
	// This function looks for a definition of localVarPath, `localVarPath := localBasePath + "path"`,
	// and replaces the `"path"` part with `link`.
	// ex: localVarPath := localBasePath + "/build-jobs/" -> `localVarPath := localBasePath + link`
	for i, stmt := range fn.Body.List {
		assign, ok := stmt.(*ast.AssignStmt)
		if !ok {
			continue
		}

		if assign.Tok == token.DEFINE && len(assign.Lhs) > 0 {
			if ident, ok := assign.Lhs[0].(*ast.Ident); ok && ident.Name == "localVarPath" {
				if expr, ok := assign.Rhs[0].(*ast.BinaryExpr); ok && len(assign.Rhs) > 0 {
					// Add empty string check before modification
					newBody := make([]ast.Stmt, 0, len(fn.Body.List)+1)
					newBody = append(newBody, fn.Body.List[:i]...)
					newBody = append(newBody, createEmptyLinkCheck())
					newBody = append(newBody, fn.Body.List[i:]...)
					fn.Body.List = newBody

					// Modify the line
					expr.Y = ast.NewIdent(resourcePath)
				}

				return
			}
		}
	}

	err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not find `localVarPath` definition")
	return
}

func addLinkQueryParamGuards(fn *ast.FuncDecl) (err error) {
	// Adds the following code before 'localVarHeaderParams' definition
	// 	linkHasOffsetParam := false
	//  linkHasLimitParam := false
	//  if parsedLink, err := url.Parse(link); err == nil {
	//    linkQuery := parsedLink.Query()
	//    linkHasOffsetParam = linkQuery.Has("offset")
	//    linkHasLimitParam = linkQuery.Has("limit")
	//  }
	insertIdx := -1
	for i, stmt := range fn.Body.List {
		assign, ok := stmt.(*ast.AssignStmt)
		if !ok || assign.Tok != token.DEFINE || len(assign.Lhs) == 0 {
			continue
		}
		if ident, ok := assign.Lhs[0].(*ast.Ident); ok && ident.Name == "localVarHeaderParams" {
			insertIdx = i
			break
		}
	}

	if insertIdx == -1 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not locate `localVarHeaderParams` definition")
		return
	}

	newStmts := []ast.Stmt{
		newBoolAssignment(pagingOffsetExistFlag),
		newBoolAssignment(pagingLimitExistFlag),
		newBoolAssignment(pagingEmbedExistFlag),
		newParsedLinkIfStmt(),
	}

	body := make([]ast.Stmt, 0, len(fn.Body.List)+len(newStmts))
	body = append(body, fn.Body.List[:insertIdx]...)
	body = append(body, newStmts...)
	body = append(body, fn.Body.List[insertIdx:]...)
	fn.Body.List = body
	return
}

func newBoolAssignment(name string) ast.Stmt {
	return &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(name)},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{ast.NewIdent("false")},
	}
}

func newParsedLinkIfStmt() ast.Stmt {
	initStmt := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("parsedLink"), ast.NewIdent("err")},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("url"),
					Sel: ast.NewIdent("Parse"),
				},
				Args: []ast.Expr{ast.NewIdent("link")},
			},
		},
	}

	cond := &ast.BinaryExpr{
		X:  ast.NewIdent("err"),
		Op: token.EQL,
		Y:  ast.NewIdent("nil"),
	}

	assignLinkQuery := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent("linkQuery")},
		Tok: token.DEFINE,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("parsedLink"),
					Sel: ast.NewIdent("Query"),
				},
			},
		},
	}

	setOffset := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(pagingOffsetExistFlag)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("linkQuery"),
					Sel: ast.NewIdent("Has"),
				},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"offset"`}},
			},
		},
	}

	setLimit := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(pagingLimitExistFlag)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("linkQuery"),
					Sel: ast.NewIdent("Has"),
				},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"limit"`}},
			},
		},
	}

	setEmbed := &ast.AssignStmt{
		Lhs: []ast.Expr{ast.NewIdent(pagingEmbedExistFlag)},
		Tok: token.ASSIGN,
		Rhs: []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("linkQuery"),
					Sel: ast.NewIdent("Has"),
				},
				Args: []ast.Expr{&ast.BasicLit{Kind: token.STRING, Value: `"embed"`}},
			},
		},
	}
	return &ast.IfStmt{
		Init: initStmt,
		Cond: cond,
		Body: &ast.BlockStmt{
			List: []ast.Stmt{
				assignLinkQuery,
				setOffset,
				setLimit,
				setEmbed,
			},
		},
	}
}

func wrapPagingParamBlocks(fn *ast.FuncDecl) {
	// Searches for if 'r.limit != nil' or 'r.offset != nil' blocks and wraps
	// these blocks in an if condition checking if the field already exists in the link
	for i, stmt := range fn.Body.List {
		if ifBlock, ok := stmt.(*ast.IfStmt); ok {
			switch {
			case isFieldNotNilCheck(ifBlock.Cond, "r", pagingLimitField):
				fn.Body.List[i] = wrapIfNotWithFlag(pagingLimitExistFlag, ifBlock)
			case isFieldNotNilCheck(ifBlock.Cond, "r", pagingOffsetField):
				fn.Body.List[i] = wrapIfNotWithFlag(pagingOffsetExistFlag, ifBlock)
			case isFieldNotNilCheck(ifBlock.Cond, "r", pagingEmbedField):
				fn.Body.List[i] = wrapIfNotWithFlag(pagingEmbedExistFlag, ifBlock)
			}
		}
	}
}

func isFieldNotNilCheck(expr ast.Expr, objectName, fieldName string) bool {
	binExpr, ok := expr.(*ast.BinaryExpr)
	if !ok || binExpr.Op != token.NEQ {
		return false
	}

	left, ok := binExpr.X.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	ident, ok := left.X.(*ast.Ident)
	if !ok || ident.Name != objectName {
		return false
	}

	if left.Sel == nil || left.Sel.Name != fieldName {
		return false
	}

	right, ok := binExpr.Y.(*ast.Ident)
	return ok && right.Name == "nil"
}

func wrapIfNotWithFlag(flagName string, stmt ast.Stmt) ast.Stmt {
	return &ast.IfStmt{
		Cond: &ast.UnaryExpr{
			Op: token.NOT,
			X:  ast.NewIdent(flagName),
		},
		Body: &ast.BlockStmt{
			List: []ast.Stmt{stmt},
		},
	}
}

func addFuncDocs(fn *ast.FuncDecl) {
	fn.Doc = &ast.CommentGroup{
		List: []*ast.Comment{
			{
				Slash: fn.Pos() - 1,
				Text: fmt.Sprintf("// Follows a collection link. This function is based on the `%s` function, with one key difference:\n"+
					"// instead of using a fixed endpoint path, `localVarPath` is defined as the base URL + the `link` provided in the arguments.\n", fn.Name.Name),
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

func getLinkFollowers(funcNameMap map[string]string, d *Data) (followers Followers, err error) {
	if d == nil {
		err = commonerrors.Newf(commonerrors.ErrUndefined, "data must be defined")
		return
	}

	if d.ClientPackagePath == "" {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "missing client package path, provide a client package path using '-c' or '--client_path'")
		return
	}

	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedSyntax | packages.NeedCompiledGoFiles | packages.NeedFiles,
		Tests: false,
		Dir:   d.ClientPackagePath,
	}
	// 'parser.ParseDir' returns deprecated 'ast.Package', so we'll use 'packages.Load' till it gets updated to return 'packages.Package'
	pkgs, loadErr := packages.Load(cfg, ".")
	if loadErr != nil {
		err = commonerrors.Newf(commonerrors.ErrInvalidDestination, "could not load packages from '%s'", d.ClientPackagePath)
		return
	}

	if len(pkgs) != 1 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "expected exactly one package to be loaded from the client module path, got %d", len(pkgs))
		return
	}

	pkg := pkgs[0]
	if len(pkg.Errors) > 0 {
		var pkgErrors string
		for _, e := range pkg.Errors {
			pkgErrors += fmt.Sprintf("\n- %s", e.Error())
		}
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "could not parse package %s:%s", pkg.Name, pkgErrors)
		return
	}

	if len(pkg.GoFiles) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "package %s has no go files", pkg.Name)
		return
	}

	if len(pkg.CompiledGoFiles) == 0 {
		err = commonerrors.Newf(commonerrors.ErrUnexpected, "package %s has no compiled Go files, it might be incomplete or have unresolved dependencies", pkg.Name)
		return
	}

	for i, file := range pkg.Syntax {
		// Filter filenames that match "api_*.go"
		filename := pkg.CompiledGoFiles[i]
		if !strings.HasPrefix(filepath.Base(filename), "api_") {
			continue
		}

		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				if followFuncName, ok := funcNameMap[fn.Name.Name]; ok {
					// Process APIService.FollowLink function
					addFuncDocs(fn)
					fn.Name.Name = followFuncName
					addNewFuncParam(fn, "link", "string")
					modifyParamErr := modifyFirstParam(fn.Type.Params)
					if modifyParamErr != nil {
						err = modifyParamErr
						return
					}
					modifyLineErr := modifyLocalVarPath(fn)
					if modifyLineErr != nil {
						err = modifyLineErr
						return
					}
					guardsErr := addLinkQueryParamGuards(fn)
					if guardsErr != nil {
						err = guardsErr
						return
					}
					wrapPagingParamBlocks(fn)
					APIServiceFollowFuncSrc, renderErr := renderSrc(fn, pkg.Fset)
					if renderErr != nil {
						err = renderErr
						return
					}
					APIServiceReceiverType, renderErr := getReceiverType(fn)
					if renderErr != nil {
						err = renderErr
						return
					}

					// Render out parts of the Request.FollowLink funciton
					paramName, paramType, renderErr := getFirstParam(fn.Type.Params, pkg.Fset)
					if renderErr != nil {
						err = renderErr
						return
					}
					requestReturns, renderErr := renderFieldListSrc(fn.Type.Results, pkg.Fset)
					if renderErr != nil {
						err = renderErr
						return
					}

					followers = append(followers, Follower{
						APIService: APIServiceReceiverType,
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

	return
}
