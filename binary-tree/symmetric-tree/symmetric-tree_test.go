package symmetric_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isSymmetric_DFS(t *testing.T) {
	t.Run("[1,2,2,3,4,4,3]", func(t *testing.T) {
		/*
			    1
			   / \
			  2   2
			 / \ / \
			3  4 4  3
		*/

		n1 := &TreeNode{Val: 1}
		n2_1 := &TreeNode{Val: 2}
		n2_2 := &TreeNode{Val: 2}
		n3_1 := &TreeNode{Val: 3}
		n3_2 := &TreeNode{Val: 3}
		n4_1 := &TreeNode{Val: 4}
		n4_2 := &TreeNode{Val: 4}

		n1.Left, n1.Right = n2_1, n2_2
		n2_1.Left, n2_1.Right = n3_1, n4_1
		n2_2.Left, n2_2.Right = n4_2, n3_2

		assert.True(t, isSymmetric_DFS(n1))
	})
}

func Test_isSymmetric_BFS(t *testing.T) {
	t.Run("[1,2,2,3,4,4,3]", func(t *testing.T) {
		/*
			    1
			   / \
			  2   2
			 / \ / \
			3  4 4  3
		*/

		n1 := &TreeNode{Val: 1}
		n2_1 := &TreeNode{Val: 2}
		n2_2 := &TreeNode{Val: 2}
		n3_1 := &TreeNode{Val: 3}
		n3_2 := &TreeNode{Val: 3}
		n4_1 := &TreeNode{Val: 4}
		n4_2 := &TreeNode{Val: 4}

		n1.Left, n1.Right = n2_1, n2_2
		n2_1.Left, n2_1.Right = n3_1, n4_1
		n2_2.Left, n2_2.Right = n4_2, n3_2

		assert.True(t, isSymmetric_BFS(n1))
	})
}
