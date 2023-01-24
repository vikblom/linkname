package main

import (
	"fmt"

	"github.com/vikblom/linkname/test/one"
)

func main() {
	// go list -f '{{ join .Deps "\n" }}' time/

	// Directive in upper, definition in lower:
	// This is the easy case: a body-less func with a linkname next to it.
	//
	// See:
	// src/time/time.go
	// time.runtimeNano

	// Directive in lower, definition in lower:
	// This is the hard case, a body-less func, no idea where the directive could be.
	// Assume in some dep?
	//
	// See:
	// src/runtime/mheap.go
	// runtime.runtime_debug_freeOSMemory
	// and
	// src/runtime/debug/stubs.go
	// debug.freeOSMemory

	fmt.Printf("link here,  define there: %v\n", one.DeclHere())
	fmt.Printf("link there, define there: %v\n", one.DeclThere())
	fmt.Printf("link here,  define here: %v\n", one.DefHere())
	fmt.Printf("link there, define there: %v\n", one.DefThere())
}
