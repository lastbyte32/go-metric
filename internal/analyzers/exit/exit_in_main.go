// Package exit -  анализатор, который ищет вызов функции os.Exit в функции main.
package exit

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "osExitAnalyzer",
	Doc:  "Analyzer for os exit in main func",
	Run:  run,
}

func run(p *analysis.Pass) (any, error) {
	for _, file := range p.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			if !isMainFunction(file, node) {
				return true
			}

			checkOSExitCalls(p, file, node.(*ast.FuncDecl).Body.List)
			return true
		})
	}
	return nil, nil
}

func isMainFunction(file *ast.File, node ast.Node) bool {
	fun, ok := node.(*ast.FuncDecl)
	if !ok {
		return false
	}

	return file.Name.Name == "main" && fun.Name.Name == "main"
}

func checkOSExitCalls(p *analysis.Pass, file *ast.File, stmts []ast.Stmt) {
	for _, stmt := range stmts {
		if callExpr, ok := stmt.(*ast.ExprStmt).X.(*ast.CallExpr); ok {
			if selectorExpr, ok := callExpr.Fun.(*ast.SelectorExpr); ok {
				if ident, ok := selectorExpr.X.(*ast.Ident); ok {
					if ident.Name == "os" && selectorExpr.Sel.Name == "Exit" {
						p.Reportf(file.Pos(), "calling os.Exit in main function")
					}
				}
			}
		}
	}
}
