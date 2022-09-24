/*
Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the MinStack class:

MinStack() initializes the stack object.
void push(int val) pushes the element val onto the stack.
void pop() removes the element on the top of the stack.
int top() gets the top element of the stack.
int getMin() retrieves the minimum element in the stack.
You must implement a solution with O(1) time complexity for each function.

Example 1:

Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
*/

package min_stack

type item struct {
	val     int
	min_val int
}

type MinStack struct {
	stackItems []item
}

func Constructor() MinStack {
	return MinStack{}
}

func (m *MinStack) Push(i int) {
	min := i
	if len(m.stackItems) > 0 && m.GetMin() < i {
		min = m.GetMin()
	}
	m.stackItems = append(m.stackItems, item{val: i, min_val: min})
}

func (m *MinStack) Pop() {
	l := len(m.stackItems)
	m.stackItems = m.stackItems[:l-1]
}

func (m *MinStack) Top() int {
	l := len(m.stackItems)
	return m.stackItems[l-1].val
}

func (m *MinStack) GetMin() int {
	l := len(m.stackItems)
	return m.stackItems[l-1].min_val
}
