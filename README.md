# Linkname

Find links between symbols created through `//go:linkname` directives in a Go package and its dependencies.

## Directive

See [https://pkg.go.dev/cmd/compile](https://pkg.go.dev/cmd/compile).

Note that the directive does not have to precede the code it applies to.

The directive lets a declaration re-use a definition (borrow C lingo) somewhere else.
The link can be at either the declaration or the definition.
That means when analyzing a file, even if the file has no directives, a definition could still be part of a linkname directive.

If the directive and the definition (actual func with a body) can be "here" or "there" there are 4 cases:
```
directive here,  defined here
directive there, defined here
directive here,  defined there
directive there, defined there.
```

See `github.com/vikblom/linkname/test/one` for each case written out in code.

## Example

For example, the test files link between each other:
```
viktor@linkname (main)> go run . ./test/one/
total links: 4
github.com/vikblom/linkname/test/two.two github.com/vikblom/linkname/test/one.DeclThere
github.com/vikblom/linkname/test/two.four github.com/vikblom/linkname/test/one.DefThere
github.com/vikblom/linkname/test/one.DefHere github.com/vikblom/linkname/test/two.three
github.com/vikblom/linkname/test/one.DeclHere github.com/vikblom/linkname/test/two.one
```

A main package pulls in `runtime` and similar deps with lots of links:
```
viktor@linkname (main)> go run . ./test/cmd/
total links: 173
reflect.zeroVal runtime.zeroVal
internal/bytealg.abigen_runtime_cmpstring runtime.cmpstring
runtime.reflect_chanrecv reflect.chanrecv
runtime.syscall_setenv_c syscall.setenv_c
runtime.poll_runtime_pollOpen internal/poll.runtime_pollOpen
runtime.sync_throw sync.throw
runtime.syscall_Exit syscall.Exit
...
```
