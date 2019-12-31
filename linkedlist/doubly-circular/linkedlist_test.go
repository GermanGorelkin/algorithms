package doubly_circular

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_String(t *testing.T) {
	l := NewList()

	t.Run("[]", func(t *testing.T) {
		assert.Equal(t, "[]", l.String())
	})
	t.Run("[1]", func(t *testing.T) {
		node := &Node{value: 1}
		node.prev, node.next = node, node
		l.tail, l.head = node, node

		assert.Equal(t, "[1]", l.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		node1 := &Node{value: 1}
		node2 := &Node{value: 2}

		node1.prev, node1.next = node2, node2
		node2.prev, node2.next = node1, node1
		l.tail, l.head = node2, node1

		assert.Equal(t, "[1, 2]", l.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		node1 := &Node{value: 1}
		node2 := &Node{value: 2}
		node3 := &Node{value: 3}

		node1.prev, node1.next = node3, node2
		node2.prev, node2.next = node1, node3
		node3.prev, node3.next = node2, node1
		l.tail, l.head = node3, node1

		assert.Equal(t, "[1, 2, 3]", l.String())
	})
}

func TestList_ReverseString(t *testing.T) {
	l := NewList()

	t.Run("[]", func(t *testing.T) {
		assert.Equal(t, "[]", l.ReverseString())
	})
	t.Run("[1]", func(t *testing.T) {
		node := &Node{value: 1}
		node.prev, node.next = node, node
		l.tail, l.head = node, node

		assert.Equal(t, "[1]", l.ReverseString())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		node1 := &Node{value: 1}
		node2 := &Node{value: 2}

		node1.prev, node1.next = node2, node2
		node2.prev, node2.next = node1, node1
		l.tail, l.head = node2, node1

		assert.Equal(t, "[2, 1]", l.ReverseString())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		node1 := &Node{value: 1}
		node2 := &Node{value: 2}
		node3 := &Node{value: 3}

		node1.prev, node1.next = node3, node2
		node2.prev, node2.next = node1, node3
		node3.prev, node3.next = node2, node1
		l.tail, l.head = node3, node1

		assert.Equal(t, "[3, 2, 1]", l.ReverseString())
	})
}

func TestList_AddHead(t *testing.T) {
	l := NewList()

	t.Run("[1]", func(t *testing.T) {
		l.AddHead(1)

		assert.Equal(t, 1, l.Size())
		assert.Equal(t, "[1]", l.String())
		assert.Equal(t, "[1]", l.ReverseString())
	})
	t.Run("[2, 1]", func(t *testing.T) {
		l.AddHead(2)

		assert.Equal(t, 2, l.Size())
		assert.Equal(t, "[2, 1]", l.String())
		assert.Equal(t, "[1, 2]", l.ReverseString())
	})
	t.Run("[3, 2, 1]", func(t *testing.T) {
		l.AddHead(3)

		assert.Equal(t, 3, l.Size())
		assert.Equal(t, "[3, 2, 1]", l.String())
		assert.Equal(t, "[1, 2, 3]", l.ReverseString())
	})
	t.Run("[4, 3, 2, 1]", func(t *testing.T) {
		l.AddHead(4)

		assert.Equal(t, 4, l.Size())
		assert.Equal(t, "[4, 3, 2, 1]", l.String())
		assert.Equal(t, "[1, 2, 3, 4]", l.ReverseString())
	})
}

func TestList_AddTail(t *testing.T) {
	l := NewList()

	t.Run("[1]", func(t *testing.T) {
		l.AddTail(1)

		assert.Equal(t, 1, l.Size())
		assert.Equal(t, "[1]", l.String())
		assert.Equal(t, "[1]", l.ReverseString())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		l.AddTail(2)

		assert.Equal(t, 2, l.Size())
		assert.Equal(t, "[1, 2]", l.String())
		assert.Equal(t, "[2, 1]", l.ReverseString())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		l.AddTail(3)

		assert.Equal(t, 3, l.Size())
		assert.Equal(t, "[1, 2, 3]", l.String())
		assert.Equal(t, "[3, 2, 1]", l.ReverseString())
	})
	t.Run("[1, 2, 3, 4]", func(t *testing.T) {
		l.AddTail(4)

		assert.Equal(t, 4, l.Size())
		assert.Equal(t, "[1, 2, 3, 4]", l.String())
		assert.Equal(t, "[4, 3, 2, 1]", l.ReverseString())
	})
}

func TestList_RemoveHead(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		l := NewList()
		n, err := l.RemoveHead()

		assert.Nil(t, n)
		assert.Equal(t, ErrEmptyList, err)
		assert.Equal(t, 0, l.Size())
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
	})
	t.Run("[1]", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		n, err := l.RemoveHead()

		assert.Nil(t, err)
		assert.Equal(t, 1, n.value)
		assert.Equal(t, 0, l.Size())
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)

		n, err := l.RemoveHead()

		assert.Nil(t, err)
		assert.Equal(t, 1, n.value)
		assert.Equal(t, 1, l.Size())
		assert.Equal(t, "[2]", l.String())
		assert.Equal(t, "[2]", l.ReverseString())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)

		n, err := l.RemoveHead()

		assert.Nil(t, err)
		assert.Equal(t, 1, n.value)
		assert.Equal(t, 2, l.Size())
		assert.Equal(t, "[2, 3]", l.String())
		assert.Equal(t, "[3, 2]", l.ReverseString())
	})
	t.Run("[1, 2, 3, 4]", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)
		l.AddTail(4)

		n, err := l.RemoveHead()

		assert.Nil(t, err)
		assert.Equal(t, 1, n.value)
		assert.Equal(t, 3, l.Size())
		assert.Equal(t, "[2, 3, 4]", l.String())
		assert.Equal(t, "[4, 3, 2]", l.ReverseString())
	})
}
