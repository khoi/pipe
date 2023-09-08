package funk

import (
	"reflect"
	"strings"
	"testing"

	"github.com/kr/pretty"
)

func TestCollect_withNonPrimitives(t *testing.T) {
	type UkCities struct {
		name string
	}

	type UKCityPointers struct {
		o *UkCities
	}

	twos := []*UKCityPointers{
		{o: &UkCities{name: "Litchfield"}},
		{o: &UkCities{name: "Bradford"}},
		{o: &UkCities{name: "Wells"}}, // Fewer than 6 characters
	}

	cityPointers := Collect(twos, func(p *UKCityPointers) bool { return len(p.o.name) > 6 }, func(t *UKCityPointers) UkCities { return *t.o })

	cityNames := Collect(cityPointers, func(c UkCities) bool { return len(c.name) > 6 }, func(t UkCities) string { return t.name })

	if len(cityNames) != len(cityPointers) {
		t.Fatalf(
			"Mapping %v values should've resulted in as many values, but got %v instead",
			len(cityPointers), len(cityNames))
	}

	expected := "LitchfieldBradford"
	got := strings.Join(cityNames, "")
	if got != expected {
		t.Fatalf("Expected to get %s but got %s", expected, got)
	}
}

func TestCollect_fromNilSlice(t *testing.T) {
	var nilInts []int = nil
	zeroLengthInts := []int{}
	mappedInts := Collect(nilInts, func(num int) bool { return num%2 == 0 }, func(num int) int { return num * 2 })

	if !reflect.DeepEqual(mappedInts, zeroLengthInts) {
		t.Fatalf(
			"Mapping %v should've resulted in an empty slice, but got %v instead",
			nilInts, nilInts)
	}
}

func TestCollect_withPrimitives(t *testing.T) {
	nums := []string{"one", "two", "three", "four", "five"}
	expectedLengths := []int{4, 4}

	lengths := Collect(nums,
		func(word string) bool { return len(word)%2 == 0 }, // only even length words
		func(word string) int { return len(word) })         // return their lengths

	if !reflect.DeepEqual(lengths, expectedLengths) {
		t.Fatalf("Collecting %s should have given %s but got %s", pretty.Sprint(nums), pretty.Sprint(expectedLengths), pretty.Sprint(lengths))
	}
}
