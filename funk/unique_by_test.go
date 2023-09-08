package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
	"golang.org/x/exp/slices"
)

func TestUniqueByStrings(t *testing.T) {
	identity := func(s string) string {
		return s
	}

	letters := []string{"a", "b", "b", "c", "c", "c"}

	expected := [3]string{"a", "b", "c"}

	got := UniqueBy(letters, identity)
	slices.Sort(got)
	t.Logf("got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))
	if cap(expected) != cap(got) {
		t.Fatalf("Expected to have same capacity of %v but got capacity %v", cap(expected), cap(got))
	}
	if len(expected) != len(got) {
		t.Fatalf("Expected to have same length of %v but got length %v", len(expected), len(got))
	}
	for i, exp := range expected {
		reflect.DeepEqual(exp, got[i])
	}
	if reflect.DeepEqual(expected, got) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(expected), pretty.Sprint(got))
	}
}

func TestUniqueByStructs(t *testing.T) {
	type abc struct {
		id   int
		name string
	}

	foo := abc{1, "Foo"}
	bar := abc{1, "Bar"}
	baz := abc{2, "Baz"}

	users := []*abc{
		&foo,
		&bar,
		&baz,
	}

	expected := []*abc{
		&foo,
		&baz,
	}

	got := UniqueBy(users, func(u *abc) int {
		return u.id
	})
	t.Logf("got: %s , expected %s", pretty.Sprint(got), pretty.Sprint(expected))

	if cap(expected) != cap(got) {
		t.Fatalf("Expected to have same capacity of %v but got capacity %v", cap(expected), cap(got))
	}

	if len(got) != len(expected) {
		t.Fatalf("Expected to get %d elements but got %d", len(expected), len(got))
	}

	for idx := range got {
		if got[idx] != expected[idx] {
			t.Errorf("Element idx = %d doesn't match %s vs %s", idx, got[idx].name, expected[idx].name)
		}
	}
}
