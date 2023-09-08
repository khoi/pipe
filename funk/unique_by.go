package funk

// UniqueBy returns a new slice with all duplicate values removed
// It takes a function to derive the key to compare on
// is stable, meaning the order of the elements is preserved
func UniqueBy[A any, B comparable](items []A, getKey func(A) B) []A {
	var final []A
	out := []A{}
	values := map[B]bool{}

	for _, item := range items {
		itemKey := getKey(item)
		_, isDuplicate := values[itemKey]
		if !isDuplicate {
			out = append(out, item)
			values[itemKey] = true
		}
	}

	final = make([]A, len(out))
	// To ensure minimum capacity
	copy(final, out)

	return final
}
