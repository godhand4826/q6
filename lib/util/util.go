package util

// AsRef returns a reference to the given value.
func AsRef[V any](v V) *V {
	return &v
}

// Map applies a function to each element of a slice and returns a new slice.
func Map[A any, B any](f func(A) B, arr []A) []B {
	res := make([]B, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}
	return res
}
