package one

import _ "unsafe"

func One() string {
	return "one"
}

// Here a bodyless function borrows an implementation from some other package.
//go:linkname Some github.com/vikblom/linkname/pkg/two.some
func Some() int

// But the linkname can also be in the other package, referring back to here.
// Requires a .s file in this package, so the compiler understands there could
// be non-go (i.e. outside the package) things to link.
func Other() int

// Here a bodyless function borrows an implementation from some other package.
//go:linkname solo
func solo() int {
	return 9
}