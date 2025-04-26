// mypkg
package analyzer

import (
	"fmt"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "ctxfield",
	Doc: "構造体にcontext.Contextを使っている場合はエラーを出すツール",
	Run: run,
}

func run(pass *analysis.Pass) (interface{}, error) {
	// importsからcontext.Contextの型情報を取得
	var ctxNamed *types.Named
	for _, imp := range pass.Pkg.Imports() {
		if imp.Path() == "context" {
			if obj := imp.Scope().Lookup("Context"); obj != nil {
				ctxNamed = obj.Type().(*types.Named)
			}
		}
	}

	if ctxNamed == nil {
		// パッケージがcontextをインポート
		return nil, nil
	}

	ctxIface := ctxNamed.Underlying().(*types.Interface)

	// パッケージ内の全ての名前付き型を調べる
	scope := pass.Pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		named, ok := obj.Type().(*types.Named)
		if !ok {
			continue
		}
		
		// 名前付き型の本当の中身の型を取得してstructでないならスキップする
		st, ok := named.Underlying().(*types.Struct)
		if !ok {
			continue
		}

		// 各フィールドをチェックする
		for i := 0; i < st.NumFields(); i++ {
			f := st.Field(i)
			t := f.Type()

			if types.AssignableTo(t, ctxNamed) || types.Implements(t, ctxIface) {
				fmt.Printf("(エラー):struct %sは、field %sとtype %sを持っている構造体です\n",
			 					name, f.Name(), t.String())
			}
		}
	}
	return nil, nil
}