package funk

import (
	"sort"
)

// Returns the keys of a map as a slice.
func Keys[T any, K comparable](input map[K]T) []K {
	out := make([]K, len(input))
	i := 0
	for key := range input {
		out[i] = key
		i++
	}
	return out
}

// KeysSorted returns a map's keys, ordered by a custom comparator.
//
// Usage with string keys:
//
//	var people map[string]Person
//	util.KeysSorted(people, func (a, b string) bool {
//	  return a < b
//	})
//
// Sorting time.Time keys:
//
//	var eventsByDate map[time.Time]Event
//	dates := util.KeysSorted(eventsByDate, time.Time.Before)
func KeysSorted[T any, K comparable](input map[K]T, less func(K, K) bool) []K {
	out := Keys(input)
	sort.SliceStable(out, func(i, j int) bool {
		return less(out[i], out[j])
	})
	return out
}
