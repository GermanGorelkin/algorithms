/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.



Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.
*/

package swap_nodes_in_pairs

import (
	ll "github.com/germangorelkin/algorithms/linkedlist"
)

func swapPairs(head *ll.ListNode) *ll.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	second := head.Next
	third := head.Next.Next

	second.Next = head
	head.Next = swapPairs(third)

	return second
}
