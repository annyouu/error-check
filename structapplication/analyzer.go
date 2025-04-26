package analyzer

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "structswithname",
	Doc: "パッケージ内の全ての構造体と型名を取得して表示する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for name, st := range StructsWithName(pass) {
		fmt.Printf("型名: %s\n 構造体: %s\n", name, st.String())
	}
	return nil, nil
}

func StructsWithName(pass *analysis.Pass) map[string]*types.Struct {
	structs := make(map[string]*types.Struct)
	scope := pass.Pkg.Scope()

	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		if obj == nil {
			continue
		}

		typ := obj.Type()
		named, ok := typ.(*types.Named)
		if !ok {
			continue
		}

		st, ok := named.Underlying().(*types.Struct)
		if !ok {
			continue
		}
		structs[name] = st
	}
	return structs
}