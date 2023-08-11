package heap

/*

 */

type IntMinHeap []int

func (h IntMinHeap) Len() int { return len(h) }

// I case of Max heap condition need to change
func (h IntMinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntMinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

// Push and pop use the pointer receiver because they modify the slice length
func (h *IntMinHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntMinHeap) Pop() any {
	old := *h
	l := len(old)
	x := old[l-1]     //in case of max heap, need to pick item from top of the heap
	*h = old[0 : l-1] //in case of max heap, remove the first, largest element

	return x
}
