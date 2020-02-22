package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestQueue_EnQueue(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueue([]int{})

		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1]", func(t *testing.T) {
		q := NewQueue([]int{})

		q.EnQueue(1)
		assert.Equal(t, "[1]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q := NewQueue([]int{1})

		q.EnQueue(2)
		assert.Equal(t, "[1, 2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{1, 2})

		q.EnQueue(3)
		assert.Equal(t, "[1, 2, 3]", q.String())
	})
}

func TestQueue_DeQueue(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueue([]int{})

		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1]", func(t *testing.T) {
		q := NewQueue([]int{1})

		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q := NewQueue([]int{1, 2})

		q.DeQueue()
		assert.Equal(t, "[2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{1, 2, 3})

		q.DeQueue()
		assert.Equal(t, "[2, 3]", q.String())
	})
	t.Run("3 dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{1, 2, 3})

		q.DeQueue()
		q.DeQueue()
		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
	t.Run("4 dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{1, 2, 3})

		q.DeQueue()
		q.DeQueue()
		q.DeQueue()
		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
}

func TestQueue_Front(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueue([]int{})
		assert.Equal(t, -1, q.Front())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{1, 2, 3})
		assert.Equal(t, 1, q.Front())
	})
	t.Run("enqueue [1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{})
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)

		assert.Equal(t, 1, q.Front())
	})
	t.Run("enqueue/dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{})
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)
		q.DeQueue()
		q.DeQueue()
		assert.Equal(t, 3, q.Front())
	})
}

func TestQueue_Rear(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueue([]int{})
		assert.Equal(t, -1, q.Rear())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{1, 2, 3})
		assert.Equal(t, 3, q.Rear())
	})
	t.Run("enqueue [1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{})
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)

		assert.Equal(t, 3, q.Rear())
	})
	t.Run("enqueue/dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueue([]int{})
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)
		q.DeQueue()
		q.DeQueue()

		assert.Equal(t, 3, q.Rear())
	})
}
