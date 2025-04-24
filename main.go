package main

import (
	// "analysis/methodset"
	// "analysis/fieldmethod"
	"analysis/errorimp"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// func main() {
// 	singlechecker.Main(methodset.Analyzer)
// }

func main() {
	singlechecker.Main(errorimp.Analyzer)
}

// func main() {
// 	singlechecker.Main(fieldmethod.Analyzer)
// }

