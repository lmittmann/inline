/*
Package inline provides the trivial functionality of inline if-else expressions
and inline value-to-pointer conversions.
*/
package inline

// If returns:
//     then      if condition == true
//     otherwise if condition != true
func If[T any](condition bool, then T, otherwise T) T {
	if condition {
		return then
	}
	return otherwise
}

// Ptr returns the pointer to the given value.
func Ptr[T any](value T) *T {
	return &value
}
