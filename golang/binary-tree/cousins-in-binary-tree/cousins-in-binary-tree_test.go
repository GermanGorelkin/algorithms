package cousins_in_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isCousins(t *testing.T) {
	t.Run("root = [1,2,3,4], x = 4, y = 3", func(t *testing.T) {
		/*
		    1
		   / \
		  2   3
		    /
		   4
		*/

		n1 := &TreeNode{Val: 1}
		n2 := &TreeNode{Val: 2}
		n3 := &TreeNode{Val: 3}
		n4 := &TreeNode{Val: 4}

		n1.Left, n1.Right = n2, n3
		n3.Left = n4

		assert.False(t, isCousins(n1, 4, 3))
	})
	t.Run("[1,2,3,null,4,null,5], x = 5, y = 4", func(t *testing.T) {
		/*
		    1
		   / \
		  2   3
		   \   \
		    4   5
		*/

		n1 := &TreeNode{Val: 1}
		n2 := &TreeNode{Val: 2}
		n3 := &TreeNode{Val: 3}
		n4 := &TreeNode{Val: 4}
		n5 := &TreeNode{Val: 5}

		n1.Left, n1.Right = n2, n3
		n2.Right = n4
		n3.Right = n5

		assert.True(t, isCousins(n1, 5, 4))
	})
}
