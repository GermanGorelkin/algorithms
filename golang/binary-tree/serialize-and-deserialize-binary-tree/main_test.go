package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCodec_serialize(t *testing.T) {
	t.Run("[1,2,3,null,null,4,5]", func(t *testing.T) {
		/*
			[1,2,3,null,null,4,5]

					 1
					/  \
				   2    3
					   / \
					  4	  5
		*/

		node1 := &TreeNode{Val: 1}
		node2 := &TreeNode{Val: 2}
		node3 := &TreeNode{Val: 3}
		node4 := &TreeNode{Val: 4}
		node5 := &TreeNode{Val: 5}

		node1.Left, node1.Right = node2, node3
		node3.Left, node3.Right = node4, node5

		ser := Constructor()
		got := ser.serialize(node1)

		assert.Equal(t, "[1,2,3,null,null,4,5]", got)
	})

	t.Run("[]", func(t *testing.T) {
		ser := Constructor()
		got := ser.serialize(nil)

		assert.Equal(t, "[]", got)
	})

	t.Run("[1]", func(t *testing.T) {
		node1 := &TreeNode{Val: 1}

		ser := Constructor()
		got := ser.serialize(node1)

		assert.Equal(t, "[1]", got)
	})

	t.Run("[1,2]", func(t *testing.T) {
		/*
			[1,2]
					 1
					/
				   2
		*/

		node1 := &TreeNode{Val: 1}
		node2 := &TreeNode{Val: 2}

		node1.Left = node2

		ser := Constructor()
		got := ser.serialize(node1)

		assert.Equal(t, "[1,2]", got)
	})
}

func TestCodec_deserialize(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		ser := Constructor()
		data := "[]"

		root := ser.deserialize(data)
		data2 := ser.serialize(root)

		assert.Equal(t, data, data2)
	})

	t.Run("[1]", func(t *testing.T) {
		ser := Constructor()
		data := "[1]"

		root := ser.deserialize(data)
		data2 := ser.serialize(root)

		assert.Equal(t, data, data2)
	})

	t.Run("[1,2]", func(t *testing.T) {
		ser := Constructor()
		data := "[1,2]"

		root := ser.deserialize(data)
		data2 := ser.serialize(root)

		assert.Equal(t, data, data2)
	})

	t.Run("[1,2,3,null,null,4,5]", func(t *testing.T) {
		ser := Constructor()
		data := "[1,2,3,null,null,4,5]"

		root := ser.deserialize(data)
		data2 := ser.serialize(root)

		assert.Equal(t, data, data2)
	})
}
