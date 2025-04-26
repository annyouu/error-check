// target4
package analyzer

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "structs",
	Doc: "パッケージ内の全ての構造体を取得して表示する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, st := range Structs(pass) {
		fmt.Println(st)
	}
	return nil, nil
}

func Structs(pass *analysis.Pass) []*types.Struct {
	var structs []*types.Struct
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
		structs = append(structs, st)
	}
	return structs
}