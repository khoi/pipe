package funk

// Exists is used to check whether at least one item in a collection satisfies the predicate
func Exists[A any](items []A, predicate func(A) bool) bool {
	for _, item := range items {
		if predicate(item) {
			return true
		}
	}
	return false
}
