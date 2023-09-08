package funk

// Unique returns a new slice with all duplicate values removed
// is stable, meaning the order of the elements is preserved
func Unique[A comparable](items []A) []A {
	var final []A
	out := []A{}
	values := map[A]bool{}
	for _, item := range items {
		_, isDuplicate := values[item]
		if !isDuplicate {
			out = append(out, item)
			values[item] = true
		}
	}
	final = make([]A, len(out))
	// To ensure minimum capacity
	copy(final, out)

	return final
}
