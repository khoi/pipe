package funk

// Flat flattens a slice of slices into a single slice
func Flat[A any](xs [][]A) []A {
	out := []A{}
	for _, x := range xs {
		out = append(out, x...)
	}
	return out
}
