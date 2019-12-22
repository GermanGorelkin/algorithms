package doubly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNode_Equals(t *testing.T) {
	t.Run("nil", func(t *testing.T) {
		var node *Node
		var other *Node
		assert.True(t, node.Equals(other))
	})
	t.Run("equal", func(t *testing.T) {
		node := &Node{value: 123}
		other := &Node{value: 123}
		assert.True(t, node.Equals(other))
	})
	t.Run("not equal", func(t *testing.T) {
		node := &Node{value: 123}
		other := &Node{value: 321}
		assert.False(t, node.Equals(other))
	})
}

func TestList_AddHead(t *testing.T) {
	l := NewList()

	l.AddHead(1)
	assert.Equal(t, "[1]", l.String())
	assert.Equal(t, "[1]", l.ReverseString())

	l.AddHead(2)
	assert.Equal(t, "[2, 1]", l.String())
	assert.Equal(t, "[1, 2]", l.ReverseString())

	l.AddHead(3)
	assert.Equal(t, "[3, 2, 1]", l.String())
	assert.Equal(t, "[1, 2, 3]", l.ReverseString())
}
func TestList_Size(t *testing.T) {
	l := NewList()
	assert.Equal(t, 0, l.Size())

	l.AddHead(1)
	assert.Equal(t, 1, l.Size())

	l.AddTail(2)
	assert.Equal(t, 2, l.Size())
}
func TestList_AddTail(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	assert.Equal(t, "[1]", l.String())
	assert.Equal(t, "[1]", l.ReverseString())

	l.AddTail(2)
	assert.Equal(t, "[1, 2]", l.String())
	assert.Equal(t, "[2, 1]", l.ReverseString())

	l.AddTail(3)
	assert.Equal(t, "[1, 2, 3]", l.String())
	assert.Equal(t, "[3, 2, 1]", l.ReverseString())
}
func TestList_AddAtIndex(t *testing.T) {
	t.Run("[] at 1", func(t *testing.T) {
		l := NewList()

		err := l.AddAtIndex(1, 1)
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
		assert.Equal(t, 0, l.Size())
		assert.Equal(t, ErrIndexOutOfRange, err)
	})
	t.Run("[] at 0", func(t *testing.T) {
		l := NewList()

		err := l.AddAtIndex(1, 0)
		assert.Equal(t, "[1]", l.String())
		assert.Equal(t, "[1]", l.ReverseString())
		assert.Equal(t, 1, l.Size())
		assert.Nil(t, err)
	})
	t.Run("[1] at 0", func(t *testing.T) {
		l := NewList()
		l.AddHead(1)

		err := l.AddAtIndex(2, 0)
		assert.Equal(t, "[2, 1]", l.String())
		assert.Equal(t, "[1, 2]", l.ReverseString())
		assert.Equal(t, 2, l.Size())
		assert.Nil(t, err)
	})
	t.Run("[1] at 1", func(t *testing.T) {
		l := NewList()
		l.AddHead(1)

		err := l.AddAtIndex(2, 1)
		assert.Equal(t, "[1, 2]", l.String())
		assert.Equal(t, "[2, 1]", l.ReverseString())
		assert.Equal(t, 2, l.Size())
		assert.Nil(t, err)
	})
	t.Run("[1,2,3] at 0", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)

		err := l.AddAtIndex(4, 0)
		assert.Equal(t, "[4, 1, 2, 3]", l.String())
		assert.Equal(t, "[3, 2, 1, 4]", l.ReverseString())
		assert.Equal(t, 4, l.Size())
		assert.Nil(t, err)
	})
	t.Run("[1,2,3] at 1", func(t *testing.T) {
		l := NewList()
		l.AddTail(1)
		l.AddTail(2)
		l.AddTail(3)

		err := l.AddAtIndex(4, 1)
		assert.Equal(t, "[1, 4, 2, 3]", l.String())
		assert.Equal(t, "[3, 2, 4, 1]", l.ReverseString())
		assert.Equal(t, 4, l.Size())
		assert.Nil(t, err)
	})
}
func TestList_RemoveHead(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		l := NewList()
		val, err := l.RemoveHead()
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
		assert.Equal(t, 0, val)
		assert.Equal(t, ErrEmptyList, err)
	})

	t.Run("[1]", func(t *testing.T) {
		l := NewList()
		l.AddHead(1)

		val, err := l.RemoveHead()
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
		assert.Equal(t, 1, val)
		assert.Nil(t, err)
		assert.Equal(t, 0, l.Size())
	})

	t.Run("[1, 2]", func(t *testing.T) {
		l := NewList()
		l.AddHead(2)
		l.AddHead(1)

		val, err := l.RemoveHead()
		assert.Equal(t, "[2]", l.String())
		assert.Equal(t, "[2]", l.ReverseString())
		assert.Equal(t, 1, val)
		assert.Nil(t, err)
		assert.Equal(t, 1, l.Size())
	})

	t.Run("[1, 2, 3]", func(t *testing.T) {
		l := NewList()
		l.AddHead(3)
		l.AddHead(2)
		l.AddHead(1)

		val, err := l.RemoveHead()
		assert.Equal(t, "[2, 3]", l.String())
		assert.Equal(t, "[3, 2]", l.ReverseString())
		assert.Equal(t, 1, val)
		assert.Nil(t, err)
		assert.Equal(t, 2, l.Size())
	})
}
func TestList_RemoveTail(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		l := NewList()
		val, err := l.RemoveTail()
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
		assert.Equal(t, 0, val)
		assert.Equal(t, ErrEmptyList, err)
	})

	t.Run("[1]", func(t *testing.T) {
		l := NewList()
		l.AddHead(1)

		val, err := l.RemoveTail()
		assert.Equal(t, "[]", l.String())
		assert.Equal(t, "[]", l.ReverseString())
		assert.Equal(t, 1, val)
		assert.Nil(t, err)
		assert.Equal(t, 0, l.Size())
	})

	t.Run("[1, 2]", func(t *testing.T) {
		l := NewList()
		l.AddHead(2)
		l.AddHead(1)

		val, err := l.RemoveTail()
		assert.Equal(t, "[1]", l.String())
		assert.Equal(t, "[1]", l.ReverseString())
		assert.Equal(t, 2, val)
		assert.Nil(t, err)
		assert.Equal(t, 1, l.Size())
	})

	t.Run("[1, 2, 3]", func(t *testing.T) {
		l := NewList()
		l.AddHead(3)
		l.AddHead(2)
		l.AddHead(1)

		val, err := l.RemoveTail()
		assert.Equal(t, "[1, 2]", l.String())
		assert.Equal(t, "[2, 1]", l.ReverseString())
		assert.Equal(t, 3, val)
		assert.Nil(t, err)
		assert.Equal(t, 2, l.Size())
	})
}
