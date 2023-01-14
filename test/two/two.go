package two

import _ "unsafe"

// one.DeclHere linknames to this.
func one() int {
	return 1
}

//go:linkname two github.com/vikblom/linkname/test/one.DeclThere
func two() int {
	return 2
}

// one.DefHere linknames to this.
func three() int

//go:linkname four github.com/vikblom/linkname/test/one.DefThere
func four() int
