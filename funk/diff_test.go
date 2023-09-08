package funk

import (
	"reflect"
	"testing"
)

func TestDiff(t *testing.T) {
	type args struct {
		xs []string
		ys []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Diff works as expected with non empty params",
			args{[]string{"A", "B", "C"}, []string{"C", "D", "E"}},
			[]string{"A", "B"},
		},
		{
			"Diff returns xs when ys nil",
			args{[]string{"A", "B", "C"}, nil},
			[]string{"A", "B", "C"},
		},
		{
			"Diff returns  nil when xs are nil",
			args{nil, []string{"C", "D", "E"}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Diff(tt.args.xs, tt.args.ys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Diff() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersect(t *testing.T) {
	type args struct {
		xs []string
		ys []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"Intersect works as expected with non empty params",
			args{[]string{"A", "B", "C"}, []string{"C", "D", "E"}},
			[]string{"C"},
		},
		{
			"Intersect returns xs when xs are the same as ys",
			args{[]string{"A", "B", "C"}, []string{"A", "B", "C"}},
			[]string{"A", "B", "C"},
		},
		{
			"Diff returns nil when ys nil",
			args{[]string{"A", "B", "C"}, nil},
			nil,
		},
		{
			"Diff returns nil when xs are nil",
			args{nil, []string{"C", "D", "E"}},
			nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Intersect(tt.args.xs, tt.args.ys); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Intersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntersectKeyMap(t *testing.T) {
	type args struct {
		m1 map[string]string
		m2 map[string]string
	}
	tests := []struct {
		name string
		args args
		want map[string]string
	}{
		{
			name: "return empty map with no shared keys",
			args: args{
				map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
				map[string]string{
					"d": "D",
					"e": "E",
					"f": "F",
				},
			},
			want: map[string]string{},
		},
		{
			name: "works as expected with some shared keys",
			args: args{
				map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
				map[string]string{
					"d": "D",
					"e": "E",
					"b": "B",
					"c": "C",
					"f": "F",
				},
			},
			want: map[string]string{
				"b": "B",
				"c": "C",
			},
		},
		{
			name: "Intersect returns same as m1 if map are identical",
			args: args{
				map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
				map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
			},
			want: map[string]string{
				"a": "A",
				"b": "B",
				"c": "C",
			},
		},
		{
			name: "returns empty map when m2 is nil",
			args: args{
				map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
				nil,
			},
			want: map[string]string{},
		},
		{
			name: "returns empty map when m1 is nil",
			args: args{
				nil,
				map[string]string{
					"a": "A",
					"b": "B",
					"c": "C",
				},
			},
			want: map[string]string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntersectKeyMap(tt.args.m1, tt.args.m2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntersectKeyMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIdenticalSets(t *testing.T) {
	type args struct {
		xs []string
		ys []string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"Diff works as expected with identical strings",
			args{[]string{"A", "B", "C"}, []string{"A", "B", "C"}},
			true,
		},
		{
			"Diff works as expected with identical sets of strings even in deifferent order",
			args{[]string{"A", "B", "C"}, []string{"C", "B", "A"}},
			true,
		},
		{
			"Diff returns false with sets of different length",
			args{[]string{"A", "A,", "B", "C"}, []string{"A", "B", "C"}},
			false,
		},
		{
			"Diff returns false with sets of same length but with different occurrences of each element",
			args{[]string{"A", "A,", "B", "C"}, []string{"A", "B", "B", "C"}},
			false,
		},
		{
			"Diff returns xs when ys nil",
			args{[]string{"A", "B", "C"}, nil},
			false,
		},
		{
			"Diff returns true when xs and ys are nil",
			args{nil, nil},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IdenticalSets(tt.args.xs, tt.args.ys); got != tt.want {
				t.Errorf("IdenticalSets() = %v, want %v", got, tt.want)
			}
		})
	}
}
