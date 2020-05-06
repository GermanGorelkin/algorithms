package search_in_a_binary_search_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_searchBST(t *testing.T) {
	t.Run("[3,9,20,null,null,15,21]", func(t *testing.T) {
		/*
		    3
		   / \
		  9  20
		    /  \
		   15   21
		*/

		n3 := &TreeNode{Val: 3}
		n9 := &TreeNode{Val: 9}
		n20 := &TreeNode{Val: 20}
		n15 := &TreeNode{Val: 15}
		n21 := &TreeNode{Val: 21}

		n3.Left, n3.Right = n9, n20
		n20.Left, n20.Right = n15, n21

		assert.Equal(t, n15, searchBST(n3, n15.Val))
	})
}
