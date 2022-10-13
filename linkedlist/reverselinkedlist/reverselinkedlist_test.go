package reverselinkedlist

import "testing"

func TestReverselinkedlist(t *testing.T) {
	llistInput := ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}
	llitsExpectedOutput := ListNode{5, &ListNode{4, &ListNode{3, &ListNode{2, &ListNode{1, nil}}}}}

	out := reverseList(&llistInput)

	if !compareLinkedList(&llitsExpectedOutput, out) {
		t.Errorf("Expected reverse linkedlist: %v, got: %v", llitsExpectedOutput, out)
	}
}

func compareLinkedList(one *ListNode, two *ListNode) bool {

	if one == nil && two == nil {
		return true
	} else if one != nil && two != nil {
		for one != nil && two != nil {
			if one.val != two.val {
				return false
			}
			one = one.next
			two = two.next
		}
		return true
	} else {
		return false
	}
}
