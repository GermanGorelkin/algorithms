package v2

import (
	ll "github.com/germangorelkin/algorithms/golang/linkedlist"
)

func swapPairs(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := &ll.ListNode{}
	dummy.Next = head
	cur := dummy

	for cur.Next != nil && cur.Next.Next != nil {
		first := cur.Next
		sec := cur.Next.Next

		cur.Next = sec

		first.Next = sec.Next
		sec.Next = first

		cur = cur.Next.Next
	}

	return dummy.Next
}
