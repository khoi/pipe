package funk

// Reduce used to collapse a slice down to a single
func Reduce[A any, B any](arr []A, f func(B, A) B, initialValue B) B {
	acc := initialValue

	for _, item := range arr {
		acc = f(acc, item)
	}

	return acc
}
