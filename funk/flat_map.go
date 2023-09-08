package funk

// FlatMap is a function that takes a slice and a function that returns a slice
// and returns a single slice with the results of the function applied to each
// element of the input slice.
// It's similar to Map but flattens the resultant array.
// It is the equivalent of calling Flat(Map(...)) but slightly more efficient.
func FlatMap[A any, From any](from []From, mapper func(From) []A) []A {
	out := []A{}
	for _, item := range from {
		out = append(out, mapper(item)...)
	}
	return out
}
