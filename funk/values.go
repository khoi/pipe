package funk

// Values returns a slice of values from a map
func Values[T any, K comparable](input map[K]T) []T {
	out := make([]T, len(input))
	i := 0
	for _, value := range input {
		out[i] = value
		i++
	}
	return out
}
