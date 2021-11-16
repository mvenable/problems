package my_strings

func RotateRight[T any](array []T, shift int) []T {
	var out []T
	for i := 0; i < len(array); i++ {
		out = append(out, array[(i + len(array) - shift) % len(array)])
	}
	return out
}

func RotateLeft[T any](array []T, shift int) []T {
	var out []T
	for i := 0; i < len(array); i++ {
		out = append(out, array[(i+shift) % len(array)])
	}
	return out
}