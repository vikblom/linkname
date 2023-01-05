// Package test is an intended target of the analyzer.
package test

import _ "unsafe"

// This is an *ast.FuncDecl w/o Body.
//go:linkname this somewhere.else
func this() int

// This is an *ast.FuncDecl w/ Body.
//go:linkname other github.com/vikblom/linkname/pkg/one.Other
func other() int {
	return 4
}
