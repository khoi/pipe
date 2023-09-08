package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"golang.org/x/exp/slices"
)

func TestUniqueStrings(t *testing.T) {
	letters := []string{"a", "b", "b", "c", "c", "c"}

	expected := [3]string{"a", "b", "c"}

	got := Unique(letters)
	slices.Sort(got)
	t.Logf("got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))
	if cap(expected) != cap(got) {
		t.Fatalf("Expected to have same capacity of %v but got capacity %v", cap(expected), cap(got))
	}
	if len(expected) != len(got) {
		t.Fatalf("Expected to have same length of %v but got length %v", len(expected), len(got))
	}
	for i, exp := range expected {
		if !reflect.DeepEqual(exp, got[i]) {
			t.Fatalf("Values for element %d did not match.  Got's [%s], vs expected's [%s]", i, pretty.Sprint(got[i]), pretty.Sprint(exp))
		}
	}
	if reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(expected), pretty.Sprint(got))
	}
}

func TestUniqueStructs(t *testing.T) {
	type abc struct {
		n string
	}

	one := abc{"a"}
	two := abc{"b"}

	letters := []*abc{
		&one,
		&two,
		&two,
	}

	expected := []*abc{
		&one,
		&two,
	}

	got := Unique(letters)
	t.Logf("got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))

	if cap(expected) != cap(got) {
		t.Fatalf("Expected to have same capacity of %v but got capacity %v", cap(expected), cap(got))
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected to get %d elements but got %d", len(expected), len(got))
	}

	for idx := range got {
		if got[idx] != expected[idx] {
			t.Errorf("Element idx = %d doesn't match %s vs %s", idx, got[idx], expected[idx])
		}
	}
}
