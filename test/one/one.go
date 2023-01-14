package one

import (
	// Ensure that two is part of the build.
	_ "github.com/vikblom/linkname/test/two"

	_ "unsafe"
)

//go:linkname DeclHere github.com/vikblom/linkname/test/two.one
func DeclHere() int

// two.two linknames to this.
func DeclThere() int

//go:linkname DefHere github.com/vikblom/linkname/test/two.three
func DefHere() int {
	return 3
}

// two.four linknames to this.
func DefThere() int {
	return 4
}
