package binary_tree_level_order_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	t.Run("[3,9,20,null,null,15,7]", func(t *testing.T) {
		/*
		    3
		   / \
		  9  20
		    /  \
		   15   7
		*/

		n3 := &TreeNode{Val: 3}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n15 := &TreeNode{Val: 15}
		n7 := &TreeNode{Val: 7}

		n3.Left, n3.Right = n9, n20
		n20.Left, n20.Right = n15, n7

		want := [][]int{
			{3},
			{9, 20},
			{15, 7},
		}

		assert.Equal(t, want, levelOrder(n3))
	})
	t.Run("[1,2,3,4,null,null,5]", func(t *testing.T) {
		/*
			    1
			   / \
			  2   3
			 /     \
			4       5
		*/

		n1 := &TreeNode{Val: 1}
		n2 := &TreeNode{Val: 2}
		n3 := &TreeNode{Val: 3}
		n4 := &TreeNode{Val: 4}
		n5 := &TreeNode{Val: 5}

		n1.Left, n1.Right = n2, n3
		n2.Left = n4
		n3.Right = n5

		want := [][]int{
			{1},
			{2, 3},
			{4, 5},
		}

		assert.Equal(t, want, levelOrder(n1))
	})
}

func Test_levelOrderDFS(t *testing.T) {
	t.Run("[3,9,20,null,null,15,7]", func(t *testing.T) {
		/*
		    3
		   / \
		  9  20
		    /  \
		   15   7
		*/

		n3 := &TreeNode{Val: 3}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n15 := &TreeNode{Val: 15}
		n7 := &TreeNode{Val: 7}

		n3.Left, n3.Right = n9, n20
		n20.Left, n20.Right = n15, n7

		want := [][]int{
			{3},
			{9, 20},
			{15, 7},
		}

		assert.Equal(t, want, levelOrderDFS(n3))
	})
	t.Run("[1,2,3,4,null,null,5]", func(t *testing.T) {
		/*
			    1
			   / \
			  2   3
			 /     \
			4       5
		*/

		n1 := &TreeNode{Val: 1}
		n2 := &TreeNode{Val: 2}
		n3 := &TreeNode{Val: 3}
		n4 := &TreeNode{Val: 4}
		n5 := &TreeNode{Val: 5}

		n1.Left, n1.Right = n2, n3
		n2.Left = n4
		n3.Right = n5

		want := [][]int{
			{1},
			{2, 3},
			{4, 5},
		}

		assert.Equal(t, want, levelOrderDFS(n1))
	})
}
