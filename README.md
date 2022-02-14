# inline [![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/inline.svg)](https://pkg.go.dev/github.com/lmittmann/inline)

Package inline provides the trivial functionality of inline if-else expressions
and inline value-to-pointer conversions.

## func [If](inline.go#L9)
```
func If[T any](condition bool, then T, otherwise T) T
```
If returns `then` if `condition` is true, otherwise `otherwise`.

## func [Ptr](inline.go#L17)
```
func  Ptr[T any](value T) *T
```
Ptr returns the pointer to the given value.
