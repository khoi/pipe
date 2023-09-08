package funk

import "github.com/kr/pretty"

// PrintMap returns a slice of strings that represent key-value pairs from a map. Used for debugging
func PrintMap[T any, K comparable](input map[K]T) []string {
	out := make([]string, len(input))
	i := 0
	for key, value := range input {
		out[i] = "key: " + pretty.Sprint(key) + "; value: " + pretty.Sprint(value)
		i++
	}
	return out
}
