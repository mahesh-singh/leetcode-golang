package threesum

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	input := []int{-2, 0, 3, -1, 4, 0, 3, 4, 1, 1, 1, -3, -5, 4, 0}
	expectedoutput := [][]int{{-5, 1, 4}, {-3, -1, 4}, {-3, 0, 3}, {-2, -1, 3}, {-2, 1, 1}, {-1, 0, 1}, {0, 0, 0}}
	output := threeSum(input)

	if !reflect.DeepEqual(expectedoutput, output) {
		t.Errorf("Expected: %v, got: %v", expectedoutput, output)
	}
}
