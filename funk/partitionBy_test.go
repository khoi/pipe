package funk

import (
	"reflect"
	"testing"

	"github.com/kr/pretty"
)

func TestPartitionBy(t *testing.T) {
	ints := []int{0, 1, 2, 3, 4}

	expectedL := []int{2, 3, 4}
	expectedR := []int{0, 1}
	partitionByL, partitionByR := PartitionBy(ints, func(i int) bool {
		return i > 1
	})

	if !reflect.DeepEqual(expectedL, partitionByL) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(expectedL), pretty.Sprint(partitionByL))
	}
	if !reflect.DeepEqual(expectedR, partitionByR) {
		t.Fatalf("Expected to get %s but got %s", pretty.Sprint(expectedR), pretty.Sprint(partitionByR))
	}
}
