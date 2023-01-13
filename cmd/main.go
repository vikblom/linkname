package main

import (
	"flag"
	"fmt"
	"go/ast"
	"log"
	"os"
	"strings"

	"golang.org/x/tools/go/packages"
)

// https://pkg.go.dev/cmd/compile

const mode packages.LoadMode = packages.NeedName |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedTypesInfo |
	packages.NeedDeps |
	packages.NeedImports

// load pattern of packages.
func load(pattern string) ([]*packages.Package, error) {
	conf := packages.Config{
		Mode:  mode,
		Tests: true,
	}
	return packages.Load(&conf, pattern)
}

// Stolen from analysisutils
func imports(file *ast.File, path string) bool {
	for _, imp := range file.Imports {
		if imp.Path.Value == path {
			return true
		}
	}
	return false
}

// linknames in f mapping a local symbol to a global symbol.
func linknames(f *ast.File) map[string]string {
	if !imports(f, `"unsafe"`) {
		return nil
	}
	links := make(map[string]string)
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
	return links
}

// bfsOverDeps applies do on each file in pkgs and pkgs deps.
func bfsOverDeps(pkgs []*packages.Package, do func(pkg string, f *ast.File)) {
	if len(pkgs) == 0 {
		return
	}
	seen := map[string]struct{}{pkgs[0].PkgPath: {}}
	for len(pkgs) > 0 {
		pkg := pkgs[0]
		pkgs = pkgs[1:]
		for path, dep := range pkg.Imports {
			if _, ok := seen[path]; !ok {
				seen[path] = struct{}{}
				pkgs = append(pkgs, dep)
			}
		}

		for _, f := range pkg.Syntax {
			do(pkg.PkgPath, f)
		}
	}
}

func main() {
	// eat '--' which is sometimes required when using go run.
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		println("usage: main ./some/pkg/or/file")
		os.Exit(1)
	}
	first := args[0]

	// BFS over imports starting from first.
	pkgs, err := load(first)
	if err != nil {
		log.Fatal(err)
	}
	// These are the links to find across all dependencies of first.
	links := make(map[string]string)
	bfsOverDeps(pkgs, func(pkg string, f *ast.File) {
		for k, v := range linknames(f) {
			// Prepend the pkg to the localname to get global refs on both sides.
			links[pkg+"."+k] = v
		}
	})

	// TODO: Should the reverse links also be added?
	// Would collide if a there is a chain of links?

	fmt.Println("total links:", len(links))
	for k, v := range links {
		fmt.Println(k, v)
	}

	// TODO: Use links to analyze a function and specify which symbols
	// are used elsewhere (defined)
	// or are implemented elsewhere (declared).
}
