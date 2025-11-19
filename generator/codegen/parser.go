package codegen

import (
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(src string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", src, 0)
}

func ParseDeclerations(src string) ([]ast.Decl, error) {
	wrapped := "package p\nfunc _(){\n" + src + "\n}\n"
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, "", wrapped, 0)
	if err != nil {
		return nil, err
	}
	return file.Decls, nil
}

func ParseStatements(src string, fset *token.FileSet) ([]ast.Stmt, error) {
	wrapped := "package p\nfunc _(){\n" + src + "\n}\n"
	file, err := parser.ParseFile(fset, "", wrapped, 0)
	if err != nil {
		return nil, err
	}
	fn := file.Decls[0].(*ast.FuncDecl)
	return fn.Body.List, nil
}

func ParseTemplateStatements(templatePath string, values any, fset *token.FileSet) ([]ast.Stmt, error) {
	rendered, err := renderTemplate(templatePath, values)
	if err != nil {
		return nil, err
	}
	return ParseStatements(rendered, fset)
}
