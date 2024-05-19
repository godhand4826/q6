package main

func AsRef[V any](v V) *V {
	return &v
}

func Map[A any, B any](f func(A) B, arr []A) []B {
	res := make([]B, len(arr))
	for i, v := range arr {
		res[i] = f(v)
	}
	return res
}
