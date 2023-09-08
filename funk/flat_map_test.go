package funk

import (
	"reflect"
	"strings"
	"testing"
)

func TestFlatMap(t *testing.T) {
	toUpper := func(xs []string) []string {
		return Map(xs, strings.ToUpper)
	}

	tests := []struct {
		name   string
		args   [][]string
		mapper func([]string) []string
		want   []string
	}{
		{
			name:   "empty",
			args:   [][]string{},
			mapper: toUpper,
			want:   []string{},
		},
		{
			name:   "all empty",
			args:   [][]string{{}, {}, {}},
			mapper: toUpper,
			want:   []string{},
		},
		{
			name:   "all non-empty",
			args:   [][]string{{"a"}, {"b", "c"}, {"d"}},
			mapper: toUpper,
			want:   []string{"A", "B", "C", "D"},
		},
		{
			name:   "some empty",
			args:   [][]string{{"a"}, {}, {"c"}},
			mapper: toUpper,
			want:   []string{"A", "C"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatMap(tt.args, tt.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v\t[%v]", got, tt.want, tt.name)
			}
		})
	}

	toLength := func(xs []string) []int {
		return Map(xs, func(x string) int { return len(x) })
	}

	transform_tests := []struct {
		name   string
		args   [][]string
		mapper func([]string) []int
		want   []int
	}{
		{
			name:   "empty",
			args:   [][]string{},
			mapper: toLength,
			want:   []int{},
		},
		{
			name:   "all empty",
			args:   [][]string{{}, {}, {}},
			mapper: toLength,
			want:   []int{},
		},
		{
			name:   "all non-empty",
			args:   [][]string{{"alpha"}, {"beta", "gamma"}, {"delta"}},
			mapper: toLength,
			want:   []int{5, 4, 5, 5},
		},
		{
			name:   "some empty",
			args:   [][]string{{"a"}, {}, {"c"}},
			mapper: toLength,
			want:   []int{1, 1},
		},
	}
	for _, tt := range transform_tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FlatMap(tt.args, tt.mapper); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FlatMap() = %v, want %v\t[%v]\t[transform_tests]", got, tt.want, tt.name)
			}
		})
	}
}
