package funk

import (
	"testing"

	"github.com/kr/pretty"
)

func TestContains(t *testing.T) {
	numbers := []int{0, 1, 2, 3}
	doesContain2 := Contains(2, numbers)
	doesNotContain4 := Contains(4, numbers)

	if !doesContain2 {
		t.Fatalf("Expected to get 'true' but got %s", pretty.Sprint(doesContain2))
	}
	if !doesNotContain4 == false {
		t.Fatalf("Expected to get 'true' but got %s", pretty.Sprint(doesNotContain4))
	}
}
