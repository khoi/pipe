package funk

// PartitionBy Splits one collection into two.
// The left collection passes the predicate the right does not
func PartitionBy[A comparable](from []A, filter func(A) bool) (left []A, right []A) {
	var l []A
	var r []A

	for _, item := range from {
		if filter(item) {
			l = append(l, item)
		} else {
			r = append(r, item)
		}
	}

	return l, r
}
