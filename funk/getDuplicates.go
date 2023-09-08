package funk

// GetDuplicates returns a new slice with all duplicate values removed
// is stable, meaning the order of the elements is preserved
func GetDuplicates[A comparable](items []A) map[A]int {
	duplicates := map[A]int{}

	values := map[A]bool{}
	for _, item := range items {
		_, isDuplicate := values[item]
		if !isDuplicate {
			values[item] = true
		} else {
			if numRepeats, notFirstRepeat := duplicates[item]; notFirstRepeat {
				duplicates[item] = numRepeats + 1
			} else {
				duplicates[item] = 1
			}
		}
	}

	return duplicates
}
