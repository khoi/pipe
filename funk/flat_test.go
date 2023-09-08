package funk

import (
	"reflect"
	"testing"
)

func TestFlat(t *testing.T) {
	tests := []struct {
		name string
		args [][]string
		want []string
	}{
		{
			name: "empty",
			args: [][]string{},
			want: []string{},
		},
		{
			name: "all empty",
			args: [][]string{{}, {}, {}},
			want: []string{},
		},
		{
			name: "all non-empty",
			args: [][]string{{"a"}, {"b", "c"}, {"d"}},
			want: []string{"a", "b", "c", "d"},
		},
		{
			name: "some empty",
			args: [][]string{{"a"}, {}, {"c"}},
			want: []string{"a", "c"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Flat(tt.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Flat() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}
