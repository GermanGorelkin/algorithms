package binary_search_tree_iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBSTIterator_Next(t *testing.T) {
	t.Run("base", func(t *testing.T) {
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

		iter := Constructor(n7)

		assert.Equal(t, 3, iter.Next())
		assert.Equal(t, 7, iter.Next())
		assert.True(t, iter.HasNext())
		assert.Equal(t, 9, iter.Next())
		assert.True(t, iter.HasNext())
		assert.Equal(t, 15, iter.Next())
		assert.True(t, iter.HasNext())
		assert.Equal(t, 20, iter.Next())
		assert.False(t, iter.HasNext())
	})

	/*
		["BSTIterator","hasNext","next","hasNext","next","hasNext","next","hasNext","next","hasNext"]
		[[[3,2,4,1]],[null],[null],[null],[null],[null],[null],[null],[null],[null]]
	*/

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

		iter := Constructor(n3)

		assert.Equal(t, 1, iter.Next())
		assert.True(t, iter.HasNext())
		assert.Equal(t, 2, iter.Next())
		assert.True(t, iter.HasNext())
		assert.Equal(t, 3, iter.Next())
		assert.True(t, iter.HasNext())
		assert.Equal(t, 4, iter.Next())
		assert.False(t, iter.HasNext())
	})
}
