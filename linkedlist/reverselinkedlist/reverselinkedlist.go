package reverselinkedlist

/*
206. Reverse Linked List
https://leetcode.com/problems/reverse-linked-list/
Given the head of a singly linked list, reverse the list, and return the reversed list.

Example 1:
Input: head = [1,2,3,4,5]
Output: [5,4,3,2,1]

Example 2:
Input: head = [1,2]
Output: [2,1]

Example 3:
Input: head = []
Output: []
*/

type ListNode struct {
	val  int
	next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode //previous node will be nil at start
	curr := head
	for curr != nil {
		//so that it will not lost when we move the direction
		tempNode := curr.next
		//curr's next will be previous node.
		curr.next = prev
		// new previous for next iteration, at the end of iteration, it will become the head
		prev = curr
		curr = tempNode // for next iteration
	}
	return prev
}
