package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// inorderTraversal inorder travels
func inorderTraversal(p *node) (s []int) {
	if p == nil {
		return s
	}

	s = append(s, inorderTraversal(p.left)...)
	s = append(s, p.val)
	s = append(s, inorderTraversal(p.right)...)

	return s
}

func Test_rightRotate(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		n10 := newNode(10)
		n20 := newNode(20)
		n30 := newNode(30)

		n20.left = n10
		n30.left = n20

		root := n30

		/*
					30
				   /
				20
			   /
			10

				20
			   /  \
			 10	   30
		*/

		newRoot := rightRotate(root)

		want := []int{10, 20, 30}
		got := inorderTraversal(newRoot)
		assert.Equal(t, want, got)

	})

	t.Run("generic", func(t *testing.T) {
		n10 := newNode(10)
		n20 := newNode(20)
		n25 := newNode(25)
		n30 := newNode(30)

		n20.left = n10
		n20.right = n25
		n30.left = n20

		root := n30

		/*
					30
				   /
				20
			   /  \
			10	   25

				   20
				 /   \
			  10	  30
			  		  /
				   	25
		*/

		newRoot := rightRotate(root)

		want := []int{10, 20, 25, 30}
		got := inorderTraversal(newRoot)
		assert.Equal(t, want, got)

	})
}

func Test_leftRotate(t *testing.T) {
	t.Run("basic", func(t *testing.T) {
		n10 := newNode(10)
		n20 := newNode(20)
		n30 := newNode(30)

		n10.right = n20
		n20.right = n30

		root := n10

		/*
			10
			  \
			   20
			    \
			     30


					20
				   /  \
				 10	   30
		*/

		newRoot := leftRotate(root)

		want := []int{10, 20, 30}
		got := inorderTraversal(newRoot)
		assert.Equal(t, want, got)

	})

	t.Run("generic", func(t *testing.T) {
		n10 := newNode(10)
		n20 := newNode(20)
		n25 := newNode(25)
		n30 := newNode(30)

		n10.right = n20
		n20.right = n30
		n20.left = n25

		root := n10

		/*
			  10
			   \
			    20
			   /  \
			  25  30

					   20
					 /   \
				  10	  30
				  	\
					25
		*/

		newRoot := leftRotate(root)

		want := []int{10, 25, 20, 30}
		got := inorderTraversal(newRoot)
		assert.Equal(t, want, got)

	})
}
