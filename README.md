# Linkname

Find links between symbols created through `//go:linkname` directives in a Go package and its dependencies.

## Directive

See [https://pkg.go.dev/cmd/compile](https://pkg.go.dev/cmd/compile).

Note that the directive does not have to precede the code it applies to.

## Example

For example, the test files link between each other:
```
viktor@linkname (main)> go run . ./test/one/
total links: 2
github.com/vikblom/linkname/test/one.solo github.com/vikblom/linkname/test/one.Solo
github.com/vikblom/linkname/test/one.Some github.com/vikblom/linkname/test/two.some
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
