package funk

// TakeMap returns a map with the n key value pairs
func TakeMap[T any, K comparable](input map[K]T, n int) map[K]T {
	if n > len(input) {
		return input
	}
	out := make(map[K]T, len(input))
	i := 1
	for key, value := range input {
		if i > n {
			return out
		}
		out[key] = value
		i++

	}
	return out
}
