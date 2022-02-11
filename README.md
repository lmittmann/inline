# inline [![Go Reference](https://pkg.go.dev/badge/github.com/lmittmann/inline.svg)](https://pkg.go.dev/github.com/lmittmann/inline)

Package inline provides the trivial functionality of an inline if-else expression.

## func [If](inline.go#L9)
<pre>
func If[T any](condition <a href="https://pkg.go.dev/builtin#bool">bool</a>, then T, otherwise T) T
</pre>
If returns `then` if `condition` is true, otherwise `otherwise`.
