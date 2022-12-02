/*
Reverse a singly linked list.

Example:

Input: 1->2->3->4->5->NULL
Output: 5->4->3->2->1->NULL
Follow up:

A linked list can be reversed either iteratively or recursively. Could you implement both?
*/

package reverse_linked_list

import ll "github.com/germangorelkin/algorithms/golang/linkedlist"

func reverseList(head *ll.ListNode) *ll.ListNode {
	if head == nil {
		return head
	}

	cur := head
	var prev, next *ll.ListNode

	for cur != nil {
		next = cur.Next
		cur.Next = prev

		prev = cur
		cur = next
	}

	return prev
}
