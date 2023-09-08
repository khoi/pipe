package funk

import (
	"reflect"
	"testing"
)

func Test_Reverse(t *testing.T) {
	tests := []struct {
		name string
		xs   []int
		want []int
	}{
		{
			name: "empty",
			xs:   []int{},
			want: []int{},
		},
		{
			name: "reverse slice orders",
			xs:   []int{1, 2, 3},
			want: []int{3, 2, 1},
		},
		{
			name: "reverse slice with unordered elements",
			xs:   []int{123, 23, 4444, 0},
			want: []int{0, 4444, 23, 123},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reversed := Reverse(tt.xs)
			if !reflect.DeepEqual(reversed, tt.want) {
				t.Errorf("Reverse() = %v, want %v, from %v", reversed, tt.want, tt.xs)
			}
		})
	}
}
