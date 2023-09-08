package funk

// Contains is used to check whether an item exists in a collection
func Contains[A comparable](x A, xs []A) bool {
	for _, item := range xs {
		if x == item {
			return true
		}
	}
	return false
}
