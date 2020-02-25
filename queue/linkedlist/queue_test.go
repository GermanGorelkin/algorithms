package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_String(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueue()

		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1]", func(t *testing.T) {
		q := NewQueue()
		nd := &node{data: 1}
		q.head, q.tail = nd, nd

		assert.Equal(t, "[1]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q := NewQueue()
		nd1 := &node{data: 1}
		nd2 := &node{data: 2}
		nd1.next = nd2
		q.head, q.tail = nd1, nd2

		assert.Equal(t, "[1, 2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueue()
		nd1 := &node{data: 1}
		nd2 := &node{data: 2}
		nd3 := &node{data: 3}
		nd1.next = nd2
		nd2.next = nd3
		q.head, q.tail = nd1, nd3

		assert.Equal(t, "[1, 2, 3]", q.String())
	})
}

func TestQueue_EnQueue(t *testing.T) {
	q := NewQueue()

	t.Run("[1]", func(t *testing.T) {
		q.EnQueue(1)

		assert.Equal(t, "[1]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q.EnQueue(2)

		assert.Equal(t, "[1, 2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q.EnQueue(3)

		assert.Equal(t, "[1, 2, 3]", q.String())
	})
	t.Run("[1, 2, 3, 4]", func(t *testing.T) {
		q.EnQueue(4)

		assert.Equal(t, "[1, 2, 3, 4]", q.String())
	})
}

func TestQueue_DeQueue(t *testing.T) {
	q := NewQueue()
	nd1 := &node{data: 1}
	nd2 := &node{data: 2}
	nd3 := &node{data: 3}
	nd1.next = nd2
	nd2.next = nd3
	q.head, q.tail = nd1, nd3
	q.size = 3

	t.Run("[1, 2, 3]", func(t *testing.T) {
		res := q.DeQueue()

		assert.True(t, res)
		assert.Equal(t, 2, q.size)
		assert.Equal(t, "[2, 3]", q.String())
	})
	t.Run("[2, 3]", func(t *testing.T) {
		res := q.DeQueue()

		assert.True(t, res)
		assert.Equal(t, 1, q.size)
		assert.Equal(t, "[3]", q.String())
	})
	t.Run("[3]", func(t *testing.T) {
		res := q.DeQueue()

		assert.True(t, res)
		assert.Equal(t, 0, q.size)
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[]", func(t *testing.T) {
		res := q.DeQueue()

		assert.False(t, res)
		assert.Equal(t, 0, q.size)
		assert.Equal(t, "[]", q.String())
	})
}
