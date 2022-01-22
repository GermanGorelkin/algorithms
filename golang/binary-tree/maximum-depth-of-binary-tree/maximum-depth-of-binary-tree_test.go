package maximum_depth_of_binary_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxDepth_TopDown(t *testing.T) {
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

		want := 3

		assert.Equal(t, want, maxDepth_TopDown(n3))
	})
}

func Test_maxDepth_BottonUp(t *testing.T) {
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

		want := 3

		assert.Equal(t, want, maxDepth_BottonUp(n3))
	})
}
