package min_stack

import "testing"

func TestGetMin(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)

	min := minStack.GetMin()

	if min != -3 {
		t.Fatalf("Min expected: %d, got:%d", -3, min)
	}
	minStack.Pop()
	top := minStack.Top()
	if top != 0 {
		t.Fatalf("Top expected: %d, got: %d", 0, top)
	}
	min = minStack.GetMin()
	if min != -2 {
		t.Fatalf("Min expected: %d, got:%d", -2, min)
	}
}
