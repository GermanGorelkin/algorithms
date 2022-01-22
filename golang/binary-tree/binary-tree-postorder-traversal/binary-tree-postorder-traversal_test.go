package binary_tree_postorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_postorderTraversal(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
		   1
		    \
		     2
		    /
		   3
		*/
		node3 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 2, Left: node3}
		root := &TreeNode{Val: 1, Right: node2}

		want := []int{3, 2, 1}

		assert.Equal(t, want, postorderTraversal(root))
	})

	t.Run("2", func(t *testing.T) {
		/*
		   1
		  /
		 2
		  \
		   3
		*/
		node3 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 2, Right: node3}
		root := &TreeNode{Val: 1, Left: node2}

		want := []int{3, 2, 1}

		assert.Equal(t, want, postorderTraversal(root))
	})

	t.Run("2", func(t *testing.T) {
		/*
			  1
			/   \
			2    4
			 \  / \
			 3  5  6
			        \
			         7
		*/
		node7 := &TreeNode{Val: 7}
		node6 := &TreeNode{Val: 6, Right: node7}
		node5 := &TreeNode{Val: 5}
		node4 := &TreeNode{Val: 4, Left: node5, Right: node6}
		node3 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 2, Right: node3}
		root := &TreeNode{Val: 1, Left: node2, Right: node4}

		want := []int{3, 2, 5, 7, 6, 4, 1}

		assert.Equal(t, want, postorderTraversal(root))
	})
}

func Test_postorderTraversal_iteratively(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
		   1
		    \
		     2
		    /
		   3
		*/
		node3 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 2, Left: node3}
		root := &TreeNode{Val: 1, Right: node2}

		want := []int{3, 2, 1}

		assert.Equal(t, want, postorderTraversal_iteratively(root))
	})

	t.Run("2", func(t *testing.T) {
		/*
		   1
		  /
		 2
		  \
		   3
		*/
		node3 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 2, Right: node3}
		root := &TreeNode{Val: 1, Left: node2}

		want := []int{3, 2, 1}

		assert.Equal(t, want, postorderTraversal_iteratively(root))
	})

	t.Run("2", func(t *testing.T) {
		/*
			  1
			/   \
			2    4
			 \  / \
			 3  5  6
			        \
			         7
		*/
		node7 := &TreeNode{Val: 7}
		node6 := &TreeNode{Val: 6, Right: node7}
		node5 := &TreeNode{Val: 5}
		node4 := &TreeNode{Val: 4, Left: node5, Right: node6}
		node3 := &TreeNode{Val: 3}
		node2 := &TreeNode{Val: 2, Right: node3}
		root := &TreeNode{Val: 1, Left: node2, Right: node4}

		want := []int{3, 2, 5, 7, 6, 4, 1}

		assert.Equal(t, want, postorderTraversal_iteratively(root))
	})
}
