package funk

// OneP returns an item in the list that is the most extreme in some way.
func OneP[A comparable](from []A, replace func(A, A) bool) A {
	if len(from) < 1 {
		panic("Cannot perform OneP on an empty array")
	}
	max := from[0]
	for _, next := range from {
		if replace(max, next) {
			max = next
		}
	}
	return max
}

// OneOrNil returns the pointer to an item in the list that is the most extreme in some way or else Nil
func OneOrNil[A any](from []*A, replace func(*A, *A) bool) *A {
	if len(from) < 1 {
		return nil
	}
	max := from[0]
	for _, next := range from {
		if replace(max, next) {
			max = next
		}
	}
	return max
}

// OneRef returns a reference to the first item in a slice if one exists, otherwise nil
func OneRef[A any](from []*A) *A {
	if len(from) > 0 {
		return from[0]
	}
	return nil
}
