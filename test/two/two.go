package two

import _ "unsafe"

func Two() string {
	return "two"
}

func some() int {
	return 2
}

func some2() int {
	return 3
}

//go:linkname other github.com/vikblom/linkname/test/one.Other
func other() int {
	return 4
}