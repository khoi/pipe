package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestOneP(t *testing.T) {
	numbers := []numberClass{
		{2, "prime"},
		{3, "prime"},
		{5, "prime"},
	}

	expected := numberClass{5, "prime"}
	selectedOne := OneP(numbers, func(i numberClass, j numberClass) bool {
		return j.number > i.number
	})

	if !reflect.DeepEqual(selectedOne, expected) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(selectedOne), pretty.Sprint(expected))
	}
}

func TestOneRef(t *testing.T) {
	one := 1
	two := 2
	three := 3
	numbers := []*int{&one, &two, &three}
	expected := &one

	found := OneRef(numbers)

	if !reflect.DeepEqual(found, expected) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(found), pretty.Sprint(expected))
	}

	emptyNumbers := []*int{}
	var expectedEmpty *int

	foundEmpty := OneRef(emptyNumbers)

	if !reflect.DeepEqual(foundEmpty, expectedEmpty) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(foundEmpty), pretty.Sprint(expectedEmpty))
	}
}
