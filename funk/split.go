package funk

// `split` divide a slice into multiple subslice of
// a given maximum size
func Split[A any](from []A, size int) [][]A {
	if len(from) <= size {
		return [][]A{from}
	}

	split := [][]A{}

	chunksCount := len(from) / size
	if len(from)%size > 0 {
		chunksCount++
	}

	for i := 0; i < chunksCount; i++ {
		limit := size + i*size
		if limit > len(from) {
			limit = len(from)
		}

		split = append(split, from[i*size:limit])
	}

	return split
}

// SplitMap divides a map into multiple subslice of maps of
// a given maximum size
func SplitMap[K comparable, V any](from map[K]V, size int) []map[K]V {
	if len(from) <= size {
		return []map[K]V{from}
	}

	splitMaps := []map[K]V{}

	count := 0
	subMap := map[K]V{}
	for k, v := range from {
		if count == size {
			splitMaps = append(splitMaps, subMap)
			subMap = map[K]V{}
			count = 0
		}
		subMap[k] = v
		count++
	}
	splitMaps = append(splitMaps, subMap)

	return splitMaps
}
