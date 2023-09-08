package funk

import (
	"reflect"
	"testing"

	"github.com/rekki/go/pkg/rekki/errors"
)

// Unbox returns a slice containing all the inputs type-asserted to T
// Like any type assertion, this panics when the assertion is not possible so pas op.
//
// this lives here because i don't immediately see a need for it and i think it's too apt to become
// yet another way for us to produce panics. if it's useful it can be trivially moved out of
// _test.go.
func Unbox[T any](xs ...interface{}) []T {
	return Map(xs, func(x any) T { return x.(T) })
}

func rt[T any](t *testing.T, xs []T) {
	out := Unbox[T](Box(xs...)...)

	if !reflect.DeepEqual(xs, out) {
		t.Errorf("Unbox(Box(%v)) = %v, want %v", xs, out, xs)
	}
}

func TestBoxRoundTrip(t *testing.T) {
	rt(t, []int{1, 2, 3, 4, 5})
	rt(t, []string{"ernie", "burt", "oscar the grouch"})
	rt(t, []struct{ x int }{{1}, {2}, {3}})
	rt(t, []error{errors.New("foo"), errors.New("bar")}) // works with interface types too
	rt(t, []*int{Pointerify(1), Pointerify(2)})
}
