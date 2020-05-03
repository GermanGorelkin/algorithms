package path_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_hasPathSum(t *testing.T) {
	t.Run("[1,2,2,3,4,4,3]", func(t *testing.T) {
		/*
			      5
			     / \
			    4   8
			   /   / \
			  11  13  4
			 /  \      \
			7    2      1
		*/

		n5 := &TreeNode{Val: 5}
		n4 := &TreeNode{Val: 4}
		n8 := &TreeNode{Val: 8}
		n11 := &TreeNode{Val: 11}
		n13 := &TreeNode{Val: 13}
		n4_2 := &TreeNode{Val: 4}
		n7 := &TreeNode{Val: 7}
		n2 := &TreeNode{Val: 2}
		n1 := &TreeNode{Val: 1}

		n5.Left, n1.Right = n4, n8
		n4.Left = n11
		n11.Left, n11.Right = n7, n2
		n8.Left, n8.Right = n13, n4
		n4_2.Right = n1

		assert.True(t, hasPathSum(n5, 22))
	})
}
