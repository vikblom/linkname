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

const loadMode packages.LoadMode = packages.NeedName |
	packages.NeedTypes |
	packages.NeedSyntax |
	packages.NeedTypesInfo |
	packages.NeedDeps |
	packages.NeedImports

// linknames in f mapping a local symbol to a global symbol.
func linknames(f *ast.File) map[string]string {
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

func transitiveImports(pkg *packages.Package) map[string]bool {
	pkgs := []*packages.Package{pkg}
	seen := map[string]bool{pkgs[0].PkgPath: true}
	for len(pkgs) > 0 {
		pkg := pkgs[0]
		pkgs = pkgs[1:]
		for path, dep := range pkg.Imports {
			if _, ok := seen[path]; !ok {
				seen[path] = true
				pkgs = append(pkgs, dep)
			}
		}
	}
	return seen
}

type linkType struct {
	hasBody    bool
	forwardDep bool
}

func main() {
	// eat '--' which is sometimes required when using go run.
	flag.Parse()
	args := flag.Args()
	if len(args) < 1 {
		println("usage: survey [packages]")
		os.Exit(1)
	}
	patterns := args[0:]

	conf := packages.Config{
		Mode:  loadMode,
		Tests: true,
	}
	pkgs, err := packages.Load(&conf, patterns...)
	if err != nil {
		log.Fatal(err)
	}

	count := map[linkType]int{}
	for _, pkg := range pkgs {
		if _, ok := pkg.Imports["unsafe"]; !ok {
			continue
		}

		deps := transitiveImports(pkg)
		for _, file := range pkg.Syntax {
			links := linknames(file)
			if len(links) == 0 {
				continue
			}
			// Pre-populate map of which funcs have bodies.
			funs := map[string]bool{}
			for _, decl := range file.Decls {
				if v, ok := decl.(*ast.FuncDecl); ok {
					funs[v.Name.Name] = v.Body != nil
				}
			}

			for k, v := range links {
				i := strings.LastIndex(v, ".")
				if i < 0 {
					continue
				}
				count[linkType{hasBody: funs[k], forwardDep: deps[v[:i]]}] += 1
			}
		}
	}

	fmt.Printf("   body + forward: %d\n", count[linkType{true, true}])
	fmt.Printf("   body + reverse: %d\n", count[linkType{true, false}])
	fmt.Printf("no body + forward: %d\n", count[linkType{false, true}])
	fmt.Printf("no body + reverse: %d\n", count[linkType{false, false}])
}
