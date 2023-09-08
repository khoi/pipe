package funk

import (
	"reflect"

	"github.com/rekki/go/pkg/rekki/errors"
)

func ToReverseMap[K comparable, V comparable](inputMap map[K]V) (map[V]K, error) {
	values := Values(inputMap)
	if len(Unique(values)) < len(values) {
		err := errors.New("Input Map has repeated values")
		return nil, err
	}
	oldKeys := Filter(Keys(inputMap), func(key K) bool { return !isNil(inputMap[key]) })
	getValueFromKey := func(key K) K { return key }
	getKeyFromValue := func(key K) V { return inputMap[key] }

	reverseMap := ToMapBespoke(oldKeys, getKeyFromValue, getValueFromKey)
	return reverseMap, nil
}

// ToReverseArrayMap Takes a map and reverses it. If there were repeat values in the original map their respective keys are put in an array
func ToReverseArrayMap[K comparable, V comparable](inputMap map[K]V) map[V][]K {
	newMap := map[V][]K{}
	for k, v := range inputMap {
		if contents, exists := newMap[v]; exists {
			newMap[v] = append(contents, k)
		} else {
			newMap[v] = []K{k}
		}
	}
	return newMap
}

func isNil(i any) bool {
	return i == nil || reflect.ValueOf(i).IsNil()
}
