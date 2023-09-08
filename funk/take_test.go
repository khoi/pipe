package funk

import (
	"reflect"
	"testing"
)

func TestTake(t *testing.T) {
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
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "count greater than length",
			args:  []string{"a", "b", "c"},
			count: 4,
			want:  []string{"a", "b", "c"},
		},
		{
			name:  "count less than or equal to zero",
			args:  []string{"a", "b", "c"},
			count: 0,
			want:  []string{},
		},
		{
			name:  "slice is nil",
			args:  nil,
			count: 3,
			want:  nil,
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
			if got := Take(tt.args, tt.count); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Take() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}
