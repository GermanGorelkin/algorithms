package ring_buffer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_String(t *testing.T) {
	t.Run("[1] [1]", func(t *testing.T) {
		q := &queue{tail: 0, head: 0, size: 5, data: []int{1}}

		assert.Equal(t, "[1]", q.String())
	})
	t.Run("[1, 2] [1, 2]", func(t *testing.T) {
		q := &queue{tail: 1, head: 0, size: 5, data: []int{1, 2}}

		assert.Equal(t, "[1, 2]", q.String())
	})
	t.Run("[1, 2, 3] [1, 2, 3]", func(t *testing.T) {
		q := &queue{tail: 2, head: 0, size: 5, data: []int{1, 2, 3}}

		assert.Equal(t, "[1, 2, 3]", q.String())
	})

	t.Run("[1, 2, 3, 4, 5] [4, 5, 1, 2, 3]", func(t *testing.T) {
		q := &queue{tail: 1, head: 2, size: 5, data: []int{4, 5, 1, 2, 3}}

		assert.Equal(t, "[1, 2, 3, 4, 5]", q.String())
	})

}

func TestQueue_EnQueue(t *testing.T) {
	q := NewQueue(5)

	t.Run("[1]", func(t *testing.T) {
		res := q.EnQueue(1)

		assert.True(t, res)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 0, q.tail)
		assert.Equal(t, "[1]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		res := q.EnQueue(2)

		assert.True(t, res)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, "[1, 2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		res := q.EnQueue(3)

		assert.True(t, res)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 2, q.tail)
		assert.Equal(t, "[1, 2, 3]", q.String())
	})
	t.Run("[1, 2, 3, 4]", func(t *testing.T) {
		res := q.EnQueue(4)

		assert.True(t, res)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 3, q.tail)
		assert.Equal(t, "[1, 2, 3, 4]", q.String())
	})
	t.Run("[1, 2, 3, 4, 5]", func(t *testing.T) {
		res := q.EnQueue(5)

		assert.True(t, res)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 4, q.tail)
		assert.Equal(t, "[1, 2, 3, 4, 5]", q.String())
	})
	t.Run("is full", func(t *testing.T) {
		res := q.EnQueue(6)

		assert.False(t, res)
		assert.Equal(t, 0, q.head)
		assert.Equal(t, 4, q.tail)
		assert.Equal(t, "[1, 2, 3, 4, 5]", q.String())
	})

	q.data = []int{0, 0, 1, 2, 3}
	q.head = 2
	q.tail = 4
	q.size = 5
	t.Run("circular [4, 0, 1, 2, 3]", func(t *testing.T) {
		res := q.EnQueue(4)

		assert.True(t, res)
		assert.Equal(t, 2, q.head)
		assert.Equal(t, 0, q.tail)
		assert.Equal(t, "[1, 2, 3, 4]", q.String())
	})
	t.Run("circular [4, 5, 1, 2, 3]", func(t *testing.T) {
		res := q.EnQueue(5)

		assert.True(t, res)
		assert.Equal(t, 2, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, "[1, 2, 3, 4, 5]", q.String())
	})

	t.Run("circular is full", func(t *testing.T) {
		res := q.EnQueue(6)

		assert.False(t, res)
		assert.Equal(t, 2, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, "[1, 2, 3, 4, 5]", q.String())
	})
}

func TestQueue_DeQueue(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueue(5)

		res := q.DeQueue()

		assert.False(t, res)
		assert.Equal(t, -1, q.head)
		assert.Equal(t, -1, q.tail)
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1]", func(t *testing.T) {
		q := &queue{data: []int{1}, head: 0, tail: 0, size: 5}

		res := q.DeQueue()

		assert.True(t, res)
		assert.Equal(t, -1, q.head)
		assert.Equal(t, -1, q.tail)
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q := &queue{data: []int{1, 2}, head: 0, tail: 1, size: 5}

		res := q.DeQueue()

		assert.True(t, res)
		assert.Equal(t, 1, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, "[2]", q.String())
	})
	t.Run("[4, 5, 1, 2, 3]", func(t *testing.T) {
		q := &queue{data: []int{4, 5, 1, 2, 3}, head: 2, tail: 1, size: 5}

		res := q.DeQueue()

		assert.True(t, res)
		assert.Equal(t, 3, q.head)
		assert.Equal(t, 1, q.tail)
		assert.Equal(t, "[2, 3, 4, 5]", q.String())
	})
}
