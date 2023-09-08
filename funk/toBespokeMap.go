package funk

import "github.com/rekki/go/pkg/rekki/errors"

// ToMapBespoke creates a map from a slice with a key-derivation and a value-derivation function
//
// eg. I have a slice of thing structs that have a field ID and field description, and I want a map of id -> description then I can do
//
//	thingForID := ToMapBespoke(xs, func (foo Foo) string { return foo.ID },  func (foo Foo) string { return foo.description })
func ToMapBespoke[T any, K comparable, V any](input []T, key func(T) K, value func(T) V) map[K]V {
	out := make(map[K]V, len(input))
	for _, item := range input {
		out[key(item)] = value(item)
	}
	return out
}

func ToBespokeReverseMap[K comparable, V any, T comparable](inputMap map[K]V, getNewKey func(V) T) (map[T]K, error) {
	values := Map(Values(inputMap), getNewKey)
	if len(Unique(values)) < len(values) {
		err := errors.New("Input Map has repeated values")
		return nil, err
	}
	keys := Keys(inputMap)
	getValueFromKey := func(key K) K { return key }
	getKeyFromMappedValue := func(key K) T { return getNewKey(inputMap[key]) }

	reverseMap := ToMapBespoke(keys, getKeyFromMappedValue, getValueFromKey)
	return reverseMap, nil
}
