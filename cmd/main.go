package main

import (
	"fmt"
	"go/ast"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/singlechecker"
)

// https://pkg.go.dev/cmd/compile

const Doc = `TODO`

var Analyzer = &analysis.Analyzer{
	Name:             "embed",
	Doc:              Doc,
	Requires:         []*analysis.Analyzer{},
	Run:              run,
	RunDespiteErrors: true,
}

func run(pass *analysis.Pass) (interface{}, error) {

	// "localname" -> "importpath.name"
	links := make(map[string]string)

	for _, f := range pass.Files {
		fmt.Printf("file: %+v\n", f.Name.Name)
		if !Imports(f, `"unsafe"`) {
			fmt.Printf("skip\n\n")
			continue
		}

		for _, grp := range f.Comments {
			for _, com := range grp.List {
				if strings.HasPrefix(com.Text, "//go:linkname") {
					parts := strings.Split(com.Text, " ")
					if len(parts) != 3 {
						continue
					}
					local := parts[1]
					importpath := parts[2]

					links[local] = importpath

				}
			}
		}
		fmt.Printf("%+v\n", links)

		// for _, decl := range f.Decls {
		// 	//ast.Print(pass.Fset, v)
		// 	if v, ok := decl.(*ast.FuncDecl); ok {
		// 		fmt.Printf("%+v\n", v)
		// 		if v.Doc == nil {
		// 			// TODO: Could be "detached"?
		// 			// Does not even need to be above the decl.
		// 			continue
		// 		}

		// 	}
		// 	// TODO: linkname variables?
		// }

		fmt.Println()
	}

	return nil, nil
}

// TODO: Stolen from analysisutils
func Imports(file *ast.File, path string) bool {
	for _, imp := range file.Imports {
		if imp.Path.Value == path {
			return true
		}
	}
	return false
}

func main() {
	singlechecker.Main(Analyzer)
}
