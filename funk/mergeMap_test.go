package funk

import (
	"reflect"
	"testing"
)

func TestMergeMap(t *testing.T) {
	base1 := map[string]string{
		"1": "one",
		"2": "two",
	}
	other1 := map[string]string{
		"3": "three",
	}
	other2 := map[string]string{
		"5": "five",
		"6": "six",
	}

	type example struct {
		name     string
		base     map[string]string
		others   []map[string]string
		expected map[string]string
	}
	examples := []example{
		{
			name: "no others map passed",
			base: base1,
			expected: map[string]string{
				"1": "one",
				"2": "two",
			},
		},
		{
			name:   "one other map passed",
			base:   base1,
			others: []map[string]string{other1},
			expected: map[string]string{
				"1": "one",
				"2": "two",
				"3": "three",
			},
		},
		{
			name:   "two other map passed",
			base:   base1,
			others: []map[string]string{other1, other2},
			expected: map[string]string{
				"1": "one",
				"2": "two",
				"3": "three",
				"5": "five",
				"6": "six",
			},
		},
	}

	for _, ex := range examples {
		res := MergeMap(ex.base, ex.others...)

		if !reflect.DeepEqual(ex.expected, res) {
			t.Fatalf("[%s] failed. Expected [%+v] but got [%+v]", ex.name, ex.expected, res)
		}
	}
}
