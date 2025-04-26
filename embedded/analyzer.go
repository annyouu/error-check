package analyzer

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "bufferfields",
	Doc: "bytes.Bufferのフィールドと埋め込み判定を出力する",
	Run: run,
}

// 型を探す関数
func TypeOf(pass *analysis.Pass, importPath, typename string) types.Type {
	for _, imp := range pass.Pkg.Imports() {
		if imp.Path() == importPath {
			obj := imp.Scope().Lookup(typename)
			if obj != nil {
				return obj.Type()
			}
		}
	}
	return nil
}

func run(pass *analysis.Pass) (interface{}, error) {
	typ := TypeOf(pass, "bytes", "Buffer")
	if typ == nil {
		return nil, nil
	}

	st, ok := typ.Underlying().(*types.Struct)
	if !ok {
		return nil, nil
	}

	for i := 0; i < st.NumFields(); i++ {
		field := st.Field(i)
		fmt.Printf("フィールド名: %s, 埋め込み: %v\n", field.Name(), field.Embedded())
	}
	return nil, nil
}