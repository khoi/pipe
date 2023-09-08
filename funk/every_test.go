package funk

import (
	"testing"
)

func TestEvery(t *testing.T) {
	tests := []struct {
		name string
		args []string
		want bool
	}{
		{
			name: "empty",
			args: []string{},
			want: true,
		},
		{
			name: "all empty",
			args: []string{"", "", ""},
			want: false,
		},
		{
			name: "all non-empty",
			args: []string{"a", "b", "c"},
			want: true,
		},
		{
			name: "some empty",
			args: []string{"a", "", "c"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Every(tt.args, func(s string) bool { return s != "" }); got != tt.want {
				t.Errorf("Every() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}
}
