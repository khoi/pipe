package funk

// Pointerify returns a pointer to a copy of the given value.
// Generally useful for coercing constants, literals and function return values to pointers. Notably
// it is not equivalent to & in the general case, because a pointer to the space allocated for the
// passed parameter is returned.
func Pointerify[A any](a A) *A {
	return &a
}
