package analyzer

import (
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

// Analyzer is Analyzer
var Analyzer = &analysis.Analyzer{
	Name:     "goprintffuncname",
	Doc:      "Checks that printf-like functions are named with `f` at the end.",
	Run:      run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspector := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
	}

	inspector.Preorder(nodeFilter, func(node ast.Node) {
		funcDecl := node.(*ast.FuncDecl)
		params := funcDecl.Type.Params.List
		if len(params) < 2 { // [0] must be format (string), [1] must be args (...interface{})
			return
		}

		firstParamType, ok := params[len(params)-2].Type.(*ast.Ident)
		if !ok { // first param type isn't identificator so it can't be of type "string"
			return
		}

		if firstParamType.Name != "string" { // first param (format) type is not string
			return
		}

		secondParamType, ok := params[len(params)-1].Type.(*ast.Ellipsis)
		if !ok { // args are not ellipsis (...args)
			return
		}

		elementType, ok := secondParamType.Elt.(*ast.InterfaceType)
		if !ok { // args are not of interface type, but we need interface{}
			return
		}

		if elementType.Methods != nil && len(elementType.Methods.List) != 0 {
			return // has >= 1 method in interface, but we need an empty interface "interface{}"
		}

		if strings.HasSuffix(funcDecl.Name.Name, "f") {
			return
		}

		pass.Reportf(node.Pos(), "printf-like formatting function '%s' should be named '%sf'",
			funcDecl.Name.Name, funcDecl.Name.Name)
	})

	/*inspect := func(node ast.Node) bool {
		funcDecl, ok := node.(*ast.FuncDecl)
		if !ok {
			return true
		}

		params := funcDecl.Type.QueryParams.List
		if len(params) != 2 { // [0] must be format (string), [1] must be args (...interface{})
			return true
		}

		firstParamType, ok := params[0].Type.(*ast.Ident)
		if !ok { // first param type isn't identificator so it can't be of type "string"
			return true
		}

		if firstParamType.Name != "string" { // first param (format) type is not string
			return true
		}

		secondParamType, ok := params[1].Type.(*ast.Ellipsis)
		if !ok { // args are not ellipsis (...args)
			return true
		}

		elementType, ok := secondParamType.Elt.(*ast.InterfaceType)
		if !ok { // args are not of interface type, but we need interface{}
			return true
		}

		if elementType.Methods != nil && len(elementType.Methods.List) != 0 {
			return true // has >= 1 method in interface, but we need an empty interface "interface{}"
		}

		if strings.HasSuffix(funcDecl.Name.Name, "f") {
			return true
		}

		pass.Reportf(node.Pos(), "printf-like formatting function '%s' should be named '%sf'",
			funcDecl.Name.Name, funcDecl.Name.Name)
		return true
	}

	for _, f := range pass.Files {
		ast.Inspect(f, inspect)
	}*/
	return nil, nil
}
