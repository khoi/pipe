package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestGroupBy(t *testing.T) {
	numbers, naNumbers, primeNumbers, semiPrimeNumbers := []numberClass{
		{0, "N/A"},
		{1, "N/A"},
		{2, "prime"},
		{3, "prime"},
		{4, "semi-prime"},
		{5, "prime"},
		{6, "semi-prime"},
	},
		[]numberClass{
			{0, "N/A"},
			{1, "N/A"},
		},
		[]numberClass{
			{2, "prime"},
			{3, "prime"},
			{5, "prime"},
		},
		[]numberClass{
			{4, "semi-prime"},
			{6, "semi-prime"},
		}

	numberMap := GroupBy(numbers, func(i numberClass) string {
		return i.label
	})

	testNumbers := func(expectedNumbers []numberClass, numberType string) {
		if !reflect.DeepEqual(expectedNumbers, numberMap[numberType]) {
			t.Fatalf("Expected to get %s but got %s", pretty.Sprint(expectedNumbers), pretty.Sprint(numberMap[numberType]))
		}
	}

	testNumbers(naNumbers, "N/A")
	testNumbers(primeNumbers, "prime")
	testNumbers(semiPrimeNumbers, "semi-prime")
}
