package v2

import (
	ll "github.com/germangorelkin/algorithms/golang/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_reverseList(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
			Input: 1->2->3->4->5->NULL
			Output: 5->4->3->2->1->NULL
		*/

		l5 := &ll.ListNode{Val: 5}
		l4 := &ll.ListNode{Val: 4, Next: l5}
		l3 := &ll.ListNode{Val: 3, Next: l4}
		l2 := &ll.ListNode{Val: 2, Next: l3}
		head := &ll.ListNode{Val: 1, Next: l2}

		newHead := reverseList(head)

		got := ll.Print(newHead)

		assert.Equal(t, "[5, 4, 3, 2, 1]", got)
	})
}
