package validate_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
		    2
		   / \
		  1   3
		*/
		n2 := &TreeNode{Val: 2}
		n1 := &TreeNode{Val: 1}
		n3 := &TreeNode{Val: 3}

		n2.Left, n2.Right = n1, n3

		assert.True(t, isValidBST(n2))
	})

	t.Run("2", func(t *testing.T) {
		/*
		    5
		   / \
		  1   4
		     / \
		    3   6
		*/
		n5 := &TreeNode{Val: 5}
		n1 := &TreeNode{Val: 1}
		n4 := &TreeNode{Val: 4}
		n3 := &TreeNode{Val: 3}
		n6 := &TreeNode{Val: 6}

		n5.Left, n5.Right = n1, n4
		n4.Left, n4.Right = n3, n6

		assert.False(t, isValidBST(n5))
	})

	t.Run("[10,5,15,null,null,6,20]", func(t *testing.T) {
		/*
		    10
		   / \
		  5  15
		     / \
		    6   20
		*/
		n10 := &TreeNode{Val: 10}
		n5 := &TreeNode{Val: 5}
		n15 := &TreeNode{Val: 15}
		n6 := &TreeNode{Val: 6}
		n20 := &TreeNode{Val: 20}

		n10.Left, n10.Right = n5, n15
		n15.Left, n15.Right = n6, n20

		assert.False(t, isValidBST(n10))
	})
}

func Test_isValidBST_iter(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
		    2
		   / \
		  1   3
		*/
		n2 := &TreeNode{Val: 2}
		n1 := &TreeNode{Val: 1}
		n3 := &TreeNode{Val: 3}

		n2.Left, n2.Right = n1, n3

		assert.True(t, isValidBST_iter(n2))
	})

	t.Run("2", func(t *testing.T) {
		/*
		    5
		   / \
		  1   4
		     / \
		    3   6
		*/
		n5 := &TreeNode{Val: 5}
		n1 := &TreeNode{Val: 1}
		n4 := &TreeNode{Val: 4}
		n3 := &TreeNode{Val: 3}
		n6 := &TreeNode{Val: 6}

		n5.Left, n5.Right = n1, n4
		n4.Left, n4.Right = n3, n6

		assert.False(t, isValidBST(n5))
	})

	t.Run("[10,5,15,null,null,6,20]", func(t *testing.T) {
		/*
		    10
		   / \
		  5  15
		     / \
		    6   20
		*/
		n10 := &TreeNode{Val: 10}
		n5 := &TreeNode{Val: 5}
		n15 := &TreeNode{Val: 15}
		n6 := &TreeNode{Val: 6}
		n20 := &TreeNode{Val: 20}

		n10.Left, n10.Right = n5, n15
		n15.Left, n15.Right = n6, n20

		assert.False(t, isValidBST_iter(n10))
	})
}

func Test_isValidBST_inorder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
		    2
		   / \
		  1   3
		*/
		n2 := &TreeNode{Val: 2}
		n1 := &TreeNode{Val: 1}
		n3 := &TreeNode{Val: 3}

		n2.Left, n2.Right = n1, n3

		assert.True(t, isValidBST_inorder(n2))
	})

	t.Run("2", func(t *testing.T) {
		/*
		    5
		   / \
		  1   4
		     / \
		    3   6
		*/
		n5 := &TreeNode{Val: 5}
		n1 := &TreeNode{Val: 1}
		n4 := &TreeNode{Val: 4}
		n3 := &TreeNode{Val: 3}
		n6 := &TreeNode{Val: 6}

		n5.Left, n5.Right = n1, n4
		n4.Left, n4.Right = n3, n6

		assert.False(t, isValidBST_inorder(n5))
	})

	t.Run("[10,5,15,null,null,6,20]", func(t *testing.T) {
		/*
		    10
		   / \
		  5  15
		     / \
		    6   20
		*/
		n10 := &TreeNode{Val: 10}
		n5 := &TreeNode{Val: 5}
		n15 := &TreeNode{Val: 15}
		n6 := &TreeNode{Val: 6}
		n20 := &TreeNode{Val: 20}

		n10.Left, n10.Right = n5, n15
		n15.Left, n15.Right = n6, n20

		assert.False(t, isValidBST_inorder(n10))
	})

	t.Run("[3,2,4,1]", func(t *testing.T) {
		/*
			    3
			   / \
			  2   4
			 /
			1
		*/
		n3 := &TreeNode{Val: 3}
		n2 := &TreeNode{Val: 2}
		n1 := &TreeNode{Val: 1}
		n4 := &TreeNode{Val: 4}

		n3.Left, n3.Right = n2, n4
		n2.Left = n1

		assert.True(t, isValidBST_inorder(n3))
	})
}
