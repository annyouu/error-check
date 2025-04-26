package analyzer

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/analysis"

)

var Analyzer = &analysis.Analyzer{
	Name: "bufferfields",
	Doc: "bytes.Bugfferのフィールドを出力する",
	Run: run,
}

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
		fmt.Printf("%[1]T %[1]v\n", st.Field(i))
	}
	return nil, nil
}
