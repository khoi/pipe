package funk

import (
	"reflect"
	"testing"
)

func TestDrop(t *testing.T) {
	tests := []struct {
		name  string
		args  []string
		count int
		want  []string
	}{
		{
			name:  "count less than length",
			args:  []string{"a", "b", "c", "d"},
			count: 3,
			want:  []string{"d"},
		},
		{
			name:  "count greater than length",
			args:  []string{"a", "b", "c"},
			count: 4,
			want:  []string{},
		},
		{
			name:  "count less than or equal to zero",
			args:  []string{"a", "b", "c"},
			count: 0,
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "slice is empty",
			args:  []string{},
			count: 3,
			want:  []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Drop(tt.args, tt.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Drop() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}
