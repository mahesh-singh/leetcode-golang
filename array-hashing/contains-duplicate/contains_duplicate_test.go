package contains_duplicate

import (
	"testing"
)

var containsDuplicateTests = []struct {
	inMessage string
	in        []int
	out       bool
}{
	{
		inMessage: "[1, 2, 3, 1]",
		in:        []int{1, 2, 3, 1},
		out:       true,
	},
	{
		inMessage: "[1, 2, 3, 4]",
		in:        []int{1, 2, 3, 4},
		out:       false,
	},
	{
		inMessage: "[1, 1, 1, 3, 3, 4, 3, 2, 4, 2]",
		in:        []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2},
		out:       true,
	},
}

func TestContainsDuplicate(t *testing.T) {
	for _, tt := range containsDuplicateTests {
		t.Run(tt.inMessage, func(t *testing.T) {
			got := containsDuplicate(tt.in)
			if got != tt.out {
				t.Errorf("got %t, want %t, given %v", got, tt.out, tt.in)
			}
		})
	}
}
