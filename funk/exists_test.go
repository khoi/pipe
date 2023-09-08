package funk

import (
	"testing"

	"github.com/kr/pretty"
)

func TestExists(t *testing.T) {
	numbers := []numberClass{
		{0, "N/A"},
		{1, "N/A"},
		{2, "prime"},
		{3, "prime"},
		{4, "semi-prime"},
	}

	containsSemiPrimes := Exists(numbers, func(number numberClass) bool {
		return number.label == "semi-prime"
	})
	containsComposites := Exists(numbers, func(number numberClass) bool {
		return number.label == "composite"
	})

	if !containsSemiPrimes == true {
		t.Fatalf("Expected to get 'true' but got %s", pretty.Sprint(containsSemiPrimes))
	}
	if !containsComposites == false {
		t.Fatalf("Expected to get 'false' but got %s", pretty.Sprint(containsComposites))
	}
}
