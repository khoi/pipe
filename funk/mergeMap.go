package funk

func MergeMap[T comparable, K comparable](base map[K]T, others ...map[K]T) map[K]T {
	for _, other := range others {
		for k, v := range other {
			base[k] = v
		}
	}

	return base
}
