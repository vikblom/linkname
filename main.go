package main

import (
	"fmt"

	"github.com/vikblom/linkname/pkg/one"
	"github.com/vikblom/linkname/pkg/two"
)

func main() {
	// pkg/two must be part of the compilation, else this all fails.

	fmt.Printf("%v\n", one.One())
	fmt.Printf("%v\n", two.Two())
	fmt.Printf("%v\n", one.Some())
	fmt.Printf("%v\n", one.Other())
	fmt.Printf("%v\n", one.solo())
}
