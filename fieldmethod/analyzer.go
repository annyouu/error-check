package fieldmethod

import (
	"fmt"
	"go/types"
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "fieldmethod",
	Doc: "フィールドまたはメソッドを取得する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

// lookupIdentはファイル内から対象識別子を探す関数
func lookupIdent(pass *analysis.Pass, name string) *ast.Ident {
	for id, obj := range pass.TypesInfo.Uses {
		if obj.Name() == name {
			return id
		}
	}
	return nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	obj := pass.TypesInfo.Uses[lookupIdent(pass, "ErrNotExist")]
	if obj == nil {
		fmt.Println("ErrNotExistが見つかりませんでした")
		return nil, nil
	}

	typ := obj.Type()  // *fs.PathError
	pkg := obj.Pkg()  // osパッケージ

	// Errorメソッドを取得(ポインタレシーバを含める)
	method, _, _ := types.LookupFieldOrMethod(typ, false, pkg, "Error")
	if method != nil {
		fmt.Printf("メソッドが見つかりました name: %s type: %s", method.Name(), method.Type())
	} else {
		fmt.Println("メソッドが見つかりませんでした")
	}

	return nil, nil
}

