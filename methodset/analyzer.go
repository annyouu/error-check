package methodset

import (
	"fmt"
	"go/types"
	"strings"

	"golang.org/x/tools/go/analysis"
    "golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name: "methodset",
	Doc: "メソッドセットを取得する",
	Run: run,
	Requires: []*analysis.Analyzer{inspect.Analyzer},
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		for _, imp := range file.Imports {
			path := strings.Trim(imp.Path.Value, `"`)
			if path != "bytes" {
				continue
			}

			// types.PkgNameをpass.TypesInfo.Usesから探す
			for _, obj := range pass.TypesInfo.Uses {
				if obj, ok := obj.(*types.PkgName); ok && obj.Imported().Path() == "bytes" {
					bytesPkg := obj.Imported()
					bufferObj := bytesPkg.Scope().Lookup("Buffer")
					if bufferObj == nil {
						pass.Reportf(imp.Pos(), "bytes.Bufferが見つかりません")
						return nil, nil
					}

					// Buffer型(named)から*Buffer型を生成
					named, ok := bufferObj.Type().(*types.Named)
					if !ok {
						pass.Reportf(imp.Pos(), "Bufferのnamedが見つかりません")
						return nil, nil
					}
					ptr := types.NewPointer(named)
					ms := types.NewMethodSet(ptr)

					for i := 0; i < ms.Len(); i++ {
						sel := ms.At(i)
						fn := sel.Obj().(*types.Func)
						sig := fn.Type().(*types.Signature)
						fmt.Printf(" %s %s\n", fn.Name(), sig.String())
					}
					return nil, nil
				}
			}
		}
	}
	return nil, nil
}