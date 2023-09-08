package funk

// Drop returns a new slice with the first n elements removed.
// If n is greater than the length of the slice, an empty slice is returned.
func Drop[T any](slice []T, n int) []T {
	if n >= len(slice) {
		return []T{}
	}
	return slice[n:]
}
