package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestGetDuplicateStrings(t *testing.T) {
	letters := []string{"a", "b", "b", "c", "c", "c"}

	expected := map[string]int{
		"b": 1,
		"c": 2,
	}

	got := GetDuplicates(letters)

	t.Logf("got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected to match but did not. got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))
	}
	for i, exp := range expected {
		if !reflect.DeepEqual(exp, got[i]) {
			t.Fatalf("Values for key %s did not match got [%s], but expected [%s]", i, pretty.Sprint(got[i]), pretty.Sprint(exp))
		}
	}
	if len(expected) != len(got) {
		t.Fatalf("Expected to have same length expected length %d but got length %d", len(expected), len(got))
	}
}

func TestDuplicateStructs(t *testing.T) {
	type abc struct {
		n string
	}

	one := abc{"a"}
	two := abc{"b"}
	three := abc{"c"}

	letters := []*abc{
		&one,
		&one,
		&one,
		&two,
		&two,
		&three,
	}

	expected := map[*abc]int{
		&one: 2,
		&two: 1,
	}

	got := GetDuplicates(letters)
	t.Logf("got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))

	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected to have same maps but got: %s, vs expected %s", pretty.Sprint(got), pretty.Sprint(expected))
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected to get %d elements but got %d", len(expected), len(got))
	}

	for idx := range got {
		if got[idx] != expected[idx] {
			t.Errorf("Element idx = %s doesn't match %d vs %d", idx, got[idx], expected[idx])
		}
	}
}
