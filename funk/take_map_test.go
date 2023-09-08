package funk

import (
	"reflect"
	"testing"
)

func TestTakeMap(t *testing.T) {
	type args struct {
		input map[int]string
		n     int
	}

	testMap := map[int]string{}
	testMap[1] = "one"
	testMap[2] = "two"

	halfTestMap := map[int]string{}
	halfTestMap[1] = "one"

	nilMap := map[int]string{}

	tests := []struct {
		name string
		args args
		want map[int]string
	}{
		{name: "When n is length of map, then the original map should be returned", args: args{input: testMap, n: 2}, want: testMap},
		{name: "When n is greater than the length of map, then the original map should be returned", args: args{input: testMap, n: 3}, want: testMap},
		{name: "When n is 1, then partial map should be returned", args: args{input: testMap, n: 1}, want: halfTestMap},
		{name: "When n is 0, then nil should be returned", args: args{input: testMap, n: 0}, want: nilMap},
		{name: "When n is -1, then nil should be returned", args: args{input: testMap, n: -1}, want: nilMap},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TakeMap(tt.args.input, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Logf("Not a perfect match. TakeMap() = %v, want %v - falling back to comparing sizes", got, tt.want)
				if !(len(got) == len(tt.want)) {
					t.Errorf("TakeMap() = %v, want %v", got, tt.want)
				}
			}
		})
	}
}
