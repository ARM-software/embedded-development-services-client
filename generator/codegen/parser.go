package codegen

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func ParseFile(src string) (*ast.File, error) {
	fset := token.NewFileSet()
	return parser.ParseFile(fset, "", src, 0)
}

func ParseDeclarations(src string) ([]ast.Decl, error) {
	wrapped := fmt.Sprintf("package p\nfunc _(){\n%s\n}\n", src)
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
