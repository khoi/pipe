package funk

// Every returns true if all elements in the collection satisfy the predicate
func Every[A any](xs []A, predicate func(A) bool) bool {
	for _, x := range xs {
		if !predicate(x) {
			return false
		}
	}
	return true
}
