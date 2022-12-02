package v2

import (
	ll "github.com/germangorelkin/algorithms/golang/linkedlist"
)

/*
Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
*/
func reverseList(head *ll.ListNode) *ll.ListNode {
	return reverse(head, nil)
}

func reverse(cur, prev *ll.ListNode) *ll.ListNode {
	if cur == nil {
		return prev
	}

	next := cur.Next
	cur.Next = prev

	return reverse(next, cur)
}
