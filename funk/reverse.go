package funk

func Reverse[T comparable](slice []T) []T {
	s := append([]T{}, slice...)

	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}

	return s
}
