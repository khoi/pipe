package funk

// Collect returns a slice of values applying the mapper function to each element that passes the filtering condition
// Collect is like a Filter Followed by Map
func Collect[A any, From any](from []From, filter func(From) bool, mapper func(From) A) []A {
	out := make([]A, 0, len(from))

	for _, item := range from {
		if filter(item) {
			out = append(out, mapper(item))
		}
	}
	return out
}
