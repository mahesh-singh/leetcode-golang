package arraytolinkedlist

/*
Create linked list from a given array (https://www.geeksforgeeks.org/create-linked-list-from-a-given-array/)

Given an array arr[] of size N. The task is to create linked list from the given array.
Examples:


Input : arr[]={1, 2, 3, 4, 5}
Output : 1->2->3->4->5

Input :arr[]={10, 11, 12, 13, 14}
Output : 10->11->12->13->14
*/

type node struct {
	val  int
	next *node
}

func arraytoLinkedlist(nums []int) node {
	var head node
	//insert into head
	for i := len(nums) - 1; i > 0; i-- {
		tmp := node{val: nums[i], next: &head}
		head = tmp
	}
	return head
}
