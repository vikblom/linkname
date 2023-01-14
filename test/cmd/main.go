package main

import (
	"fmt"

	"github.com/vikblom/linkname/test/one"
)

func main() {
	fmt.Printf("link here,  define there: %v\n", one.DeclHere())
	fmt.Printf("link there, define there: %v\n", one.DeclThere())
	fmt.Printf("link here,  define here: %v\n", one.DefHere())
	fmt.Printf("link there, define there: %v\n", one.DefThere())
}
