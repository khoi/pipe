package funk

import (
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

func TestMap_withNonPrimitives(t *testing.T) {
	type One struct {
		name string
	}

	type Two struct {
		o *One
	}

	twos := []*Two{
		{o: &One{name: "a"}},
		{o: &One{name: "b"}},
		{o: &One{name: "b"}},
	}

	ones := Map(twos, func(t *Two) One { return *t.o })

	strs := Map(ones, func(t One) string { return t.name })

	if len(strs) != len(ones) {
		t.Fatalf(
			"Mapping %v values should've resulted in as many values, but got %v instead",
			len(ones), len(strs))
	}

	expected := "abb"
	got := strings.Join(strs, "")
	if got != expected {
		t.Fatalf("Expected to get %s but got %s", expected, got)
	}
}

func TestMap_fromNilSlice(t *testing.T) {
	var nilInts []int = nil
	zeroLengthInts := []int{}
	mappedInts := Map(nilInts, func(num int) int { return num * 2 })

	if !reflect.DeepEqual(mappedInts, zeroLengthInts) {
		t.Fatalf(
			"Mapping %v should've resulted in an empty slice, but got %v instead",
			nilInts, nilInts)
	}
}

func TestMap_withPrimitives(t *testing.T) {
	nums := []string{"one", "two", "three"}
	expectedLengths := []int{3, 3, 5}

	lengths := Map(nums, func(word string) int {
		return len(word)
	})
	if !reflect.DeepEqual(lengths, expectedLengths) {
		t.Fatalf("Mapping %s should have given %s but got %s", pretty.Sprint(nums), pretty.Sprint(expectedLengths), pretty.Sprint(lengths))
	}
}
