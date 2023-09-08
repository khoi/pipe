package funk

import (
	"reflect"
	"testing"
)

const SIZE = 4

func TestSplit_SliceSmallerThanLength(t *testing.T) {
	from := []string{"one", "two"}
	want := [][]string{{"one", "two"}}

	got := Split(from, SIZE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split() = %v, want %v\t[Slice smaller than size]", got, want)
	}
}

func TestSplit_SliceSplitPerfectly(t *testing.T) {
	from := []int{1, 2, 3, 4, 5, 6, 7, 8}
	want := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
	}

	got := Split(from, SIZE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split() = %v, want %v\t[Slice split perfectly]", got, want)
	}
}

func TestSplit_SliceDontSplitPerfectly(t *testing.T) {
	from := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	want := [][]int{
		{1, 2, 3, 4},
		{5, 6, 7, 8},
		{9, 10},
	}

	got := Split(from, SIZE)

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Split() = %v, want %v\t[Slice don't split perfectly]", got, want)
	}
}

func TestSplitMap(t *testing.T) {
	type args struct {
		from map[int]string
		size int
	}

	testMap1 := map[int]string{1: "one", 2: "two"}
	testMap2 := map[int]string{1: "one", 2: "two", 3: "three"}
	testMap3 := map[int]string{1: "one", 2: "two", 3: "three", 4: "four", 5: "five", 6: "six", 7: "seven"}
	tests := []struct {
		name string
		args args
		want []map[int]string
	}{
		{
			"map has fewer elements than size returns an array with 1 element",
			args{from: testMap1, size: 4},
			[]map[int]string{testMap1},
		},
		{
			"map has more elements than size returns an array with more than 1 element",
			args{from: testMap1, size: 1},
			[]map[int]string{{1: "one"}, {2: "two"}},
		},
		{
			"map has more elements than size returns an array with more than 1 element of different sizes",
			args{from: testMap2, size: 2},
			[]map[int]string{testMap1, {3: "three"}},
		},
		{
			"map has more elements than size returns an array with more than 1 element of different sizes",
			args{from: testMap3, size: 2},
			[]map[int]string{testMap1, {3: "three", 4: "four"}, {5: "five", 6: "six"}, {7: "seven"}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SplitMap(tt.args.from, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Logf("Not a perfect match. TakeMap() = %v, want %v - falling back to comparing sizes", got, tt.want)
				if !(len(got) == len(tt.want)) {
					t.Errorf("SplitMap() = %v, want %v \n Should have %d sub maps but got %d", got, tt.want, len(tt.want), len(got))
				}
				if !(len(countElems(got)) == len(countElems(tt.want))) {
					t.Errorf("SplitMap() = %v, want %v \n Should have %d total elements but got %d", got, tt.want, len(countElems(tt.want)), len(countElems(got)))
				}
			}
		})
	}
}

func countElems(kv []map[int]string) []string {
	elems := []string{}
	for _, v := range kv {
		elems = append(elems, Values(v)...)
	}

	return elems
}
