package funk

import (
	"reflect"
	"testing"
)

func TestRange(t *testing.T) {
	tests := []struct {
		name  string
		xs    []int
		start int
		end   int
		want  []int
	}{
		{
			name:  "empty",
			xs:    []int{},
			start: 0,
			end:   0,
			want:  []int{},
		},
		{
			name:  "start out of range",
			xs:    []int{1, 2, 3},
			start: 4,
			end:   0,
			want:  []int{},
		},
		{
			name:  "end out of range",
			xs:    []int{1, 2, 3},
			start: 0,
			end:   4,
			want:  []int{1, 2, 3},
		},
		{
			name:  "start negative",
			xs:    []int{1, 2, 3},
			start: -1,
			end:   0,
			want:  []int{},
		},
		{
			name:  "end negative",
			xs:    []int{1, 2, 3},
			start: 0,
			end:   -1,
			want:  []int{1, 2},
		},
		{
			name:  "safe range",
			xs:    []int{1, 2, 3},
			start: 1,
			end:   2,
			want:  []int{2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Range(tt.xs, tt.start, tt.end)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Range() = %v, want %v", got, tt.want)
			}
		})
	}
}
