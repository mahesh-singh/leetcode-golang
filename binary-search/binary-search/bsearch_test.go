package bsearch

import "testing"

func TestBSearch(t *testing.T) {
	got := bsearch([]int{-1, 0, 3, 5, 9, 12}, 9)
	if got != 4 {
		t.Errorf("Expected: %v, got: %v", 4, got)
	}

	got = bsearch([]int{-1, 0, 3, 5, 9, 12}, 2)
	if got != -1 {
		t.Errorf("Expected: %v, got: %v", -1, got)
	}
}
