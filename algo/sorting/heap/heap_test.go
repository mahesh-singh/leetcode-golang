package heap

import (
	"container/heap"
	"testing"
)

func TestMinHeap(t *testing.T) {
	nums := &IntMinHeap{6, 3, 1, 7}
	heap.Init(nums)

	min := heap.Pop(nums)
	if min != 1 {
		t.Errorf("Expected: %d, got: %d", 1, min)
	}

	heap.Push(nums, 2)
	min = heap.Pop(nums)
	if min != 2 {
		t.Errorf("Expected: %d, got: %d", 2, min)
	}
}
