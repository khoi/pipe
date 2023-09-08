package funk

// Box returns a slice containing each input boxed into an interface{}
// useful for things like database/sql and friends that need interface{}.
func Box[T any](xs ...T) []any {
	return Map(xs, func(x T) any { return x })
}

// there is an Unbox implementation in box_test.go that you can move here if a real use for it
// arises.
