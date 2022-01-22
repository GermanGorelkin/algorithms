package delete_node_in_a_bst

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func preorderTraversal(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	s = append(s, root.Val)
	s = append(s, preorderTraversal(root.Left)...)
	s = append(s, preorderTraversal(root.Right)...)

	return s
}

func Test_deleteNode(t *testing.T) {
	t.Run("not found key", func(t *testing.T) {
		/*
		    7
		   / \
		  3   15
		     / \
		    9   20
		*/
		n7 := &TreeNode{Val: 7}
		n3 := &TreeNode{Val: 3}
		n15 := &TreeNode{Val: 15}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}

		n7.Left, n7.Right = n3, n15
		n15.Left, n15.Right = n9, n20

		newRoot := deleteNode(n7, 9999)
		got := preorderTraversal(newRoot)

		want := []int{7, 3, 15, 9, 20}

		assert.Equal(t, want, got)
	})
	t.Run("one node", func(t *testing.T) {
		/*
		    7
		   / \
		*/
		n7 := &TreeNode{Val: 7}

		newRoot := deleteNode(n7, 7)
		got := preorderTraversal(newRoot)

		assert.Nil(t, got)
	})
	t.Run("non children", func(t *testing.T) {
		/*
		    7
		   / \
		  3   15
		     / \
		    9   20
		*/
		n7 := &TreeNode{Val: 7}
		n3 := &TreeNode{Val: 3}
		n15 := &TreeNode{Val: 15}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}

		n7.Left, n7.Right = n3, n15
		n15.Left, n15.Right = n9, n20

		newRoot := deleteNode(n7, 3)
		got := preorderTraversal(newRoot)

		want := []int{7, 15, 9, 20}

		assert.Equal(t, want, got)
	})
	t.Run("one child", func(t *testing.T) {
		/*
			    7
			   / \
			  3   15
			     / \
			    9   20
					  \
			           21
		*/
		n7 := &TreeNode{Val: 7}
		n3 := &TreeNode{Val: 3}
		n15 := &TreeNode{Val: 15}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n21 := &TreeNode{Val: 21}

		n7.Left, n7.Right = n3, n15
		n15.Left, n15.Right = n9, n20
		n20.Right = n21

		newRoot := deleteNode(n7, 20)
		got := preorderTraversal(newRoot)

		want := []int{7, 3, 15, 9, 21}

		assert.Equal(t, want, got)
	})
	t.Run("node is root. one child", func(t *testing.T) {
		/*
				    7
				     \
			         15
				     / \
				    9   20
						  \
				           21
		*/
		n7 := &TreeNode{Val: 7}
		n15 := &TreeNode{Val: 15}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n21 := &TreeNode{Val: 21}

		n7.Left = n15
		n15.Left, n15.Right = n9, n20
		n20.Right = n21

		newRoot := deleteNode(n7, 7)
		got := preorderTraversal(newRoot)

		want := []int{15, 9, 20, 21}

		assert.Equal(t, want, got)
	})
	t.Run("node is root. two children", func(t *testing.T) {
		/*
			    7
			   / \
			  3   15
			     / \
			    9   20
					  \
			           21
		*/
		n7 := &TreeNode{Val: 7}
		n3 := &TreeNode{Val: 3}
		n15 := &TreeNode{Val: 15}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n21 := &TreeNode{Val: 21}

		n7.Left, n7.Right = n3, n15
		n15.Left, n15.Right = n9, n20
		n20.Right = n21

		newRoot := deleteNode(n7, 7)
		got := preorderTraversal(newRoot)

		want := []int{15, 9, 3, 20, 21}

		assert.Equal(t, want, got)
	})
	t.Run("two children", func(t *testing.T) {
		/*
			    7
			   / \
			  3   15
			     / \
			    9   20
					  \
			           21
		*/
		n7 := &TreeNode{Val: 7}
		n3 := &TreeNode{Val: 3}
		n15 := &TreeNode{Val: 15}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n21 := &TreeNode{Val: 21}

		n7.Left, n7.Right = n3, n15
		n15.Left, n15.Right = n9, n20
		n20.Right = n21

		newRoot := deleteNode(n7, 15)
		got := preorderTraversal(newRoot)

		want := []int{7, 3, 20, 9, 21}

		assert.Equal(t, want, got)
	})
}

func Test_search(t *testing.T) {
	/*
	    7
	   / \
	  3   15
	     / \
	    9   20
	*/
	n7 := &TreeNode{Val: 7}
	n3 := &TreeNode{Val: 3}
	n15 := &TreeNode{Val: 15}
	n9 := &TreeNode{Val: 9}
	n20 := &TreeNode{Val: 20}

	n7.Left, n7.Right = n3, n15
	n15.Left, n15.Right = n9, n20

	t.Run("1", func(t *testing.T) {
		node, parent := search(n7, 20)
		assert.Equal(t, n20, node)
		assert.Equal(t, n15, parent)
	})
	t.Run("2", func(t *testing.T) {
		node, parent := search(n7, 15)
		assert.Equal(t, n15, node)
		assert.Equal(t, n7, parent)
	})
	t.Run("3", func(t *testing.T) {
		node, parent := search(n7, 7)
		assert.Equal(t, n7, node)
		assert.Nil(t, parent)
	})
	t.Run("4", func(t *testing.T) {
		node, parent := search(n7, 3)
		assert.Equal(t, n3, node)
		assert.Equal(t, n7, parent)
	})
}
