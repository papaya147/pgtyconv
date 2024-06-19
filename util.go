package pgtycoon

func ternary[T any](condition bool, trueVal, falseVal T) T {
	if condition {
		return trueVal
	}
	return falseVal
}

func isZero[T comparable](t T) bool {
	var z T
	return t == z
}

func zero[T any]() T {
	var z T
	return z
}

func zeroPtr[T any]() *T {
	var z T
	return &z
}
