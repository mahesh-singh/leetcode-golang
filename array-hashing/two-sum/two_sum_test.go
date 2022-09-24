package two_sum

import (
	"reflect"
	"testing"
)

var twoSumTests = []struct {
	inMessags string
	in        []int
	target    int
	out       []int
}{
	{"[2,7,11,15]",
		[]int{2, 7, 11, 15},
		9,
		[]int{0, 1}},
	{
		"[3,2,4]",
		[]int{3, 2, 4},
		6,
		[]int{1, 2},
	},
	{
		"[3,3]",
		[]int{3, 3},
		6,
		[]int{0, 1},
	},
}

func TestTwoSum(t *testing.T) {
	for _, tt := range twoSumTests {
		t.Run(tt.inMessags, func(t *testing.T) {
			got := twoSum(tt.in, tt.target)
			if !reflect.DeepEqual(got, tt.out) {
				t.Errorf("got %v, want %v, given %v", got, tt.out, tt.in)
			}
		})
	}
}
