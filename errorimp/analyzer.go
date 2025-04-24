package errorimp

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "errorimp",
	Doc: "errorインタフェースを実装している型を取得する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	// errorインターフェースを取得する
	obj := types.Universe.Lookup("error")
	if obj == nil {
		return nil, fmt.Errorf("builtinが見つかりませんでした")
	}

	errorIface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok {
		return nil, fmt.Errorf("builtin errorはinterfaceではないです")
	}

	// このパッケージのスコープから型定義を列挙
	pkg := pass.Pkg
	scope := pkg.Scope()

	fmt.Printf("Package: %q\n", pkg.Path())

	for _, name := range scope.Names() {
		// 型定義(TypeName)のみ対象にする
		obj := scope.Lookup(name)
		tn, isTypeName := obj.(*types.TypeName)
		if !isTypeName {
			continue
		}

		// named型であることを確認
		named, isNamed := tn.Type().(*types.Named)
		if !isNamed {
			continue
		}

		// 値レシーバで実装しているか & ポインタレシーバで実装しているか両方チェックする
		implementsValue := types.Implements(named, errorIface)
		implementsPointer := types.Implements(types.NewPointer(named), errorIface)

		if implementsValue || implementsPointer {
			var via []string
			if implementsValue {
				via = append(via, name)
			}

			if implementsPointer {
				via = append(via, "*"+name)
			}
			fmt.Printf("name: %s, via: %v\n", name, via)
		}
	}
	return nil, nil
}