package main

import (
	"mylinter/pkg/analyzer"

	"golang.org/x/tools/go/analysis/singlechecker"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)

	// Don't use it: just to not crash on -unsafeptr flag from go vet
	//flag.Bool("unsafeptr", false, "")

}
