package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestFilter(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}

	expected := []int{2, 3, 4}
	filtered := Filter(ints, func(i int) bool {
		return i > 1
	})

	if !reflect.DeepEqual(filtered, expected) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(filtered), pretty.Sprint(expected))
	}
}

type numberClass struct {
	number int
	label  string
}
