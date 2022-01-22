package v2

import (
	ll "github.com/germangorelkin/algorithms/linkedlist"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_swapPairs(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		l4 := &ll.ListNode{Val: 4}
		l3 := &ll.ListNode{Val: 3, Next: l4}
		l2 := &ll.ListNode{Val: 2, Next: l3}
		head := &ll.ListNode{Val: 1, Next: l2}

		newHead := swapPairs(head)
		got := ll.Print(newHead)

		assert.Equal(t, "[2, 1, 4, 3]", got)
	})
	t.Run("2", func(t *testing.T) {
		l2 := &ll.ListNode{Val: 2}
		head := &ll.ListNode{Val: 1, Next: l2}

		newHead := swapPairs(head)
		got := ll.Print(newHead)

		assert.Equal(t, "[2, 1]", got)
	})
}
