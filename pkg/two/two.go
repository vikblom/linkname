package two

import _ "unsafe"

func Two() string {
	return "two"
}

func some() int  {
	return 2
}

//go:linkname other github.com/vikblom/linkname/pkg/one.Other
func other() int {
	return 4
}
