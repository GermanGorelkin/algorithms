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

func TestQueueImp_EnQueue(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)

		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)

		q.EnQueue(1)
		assert.Equal(t, "[1]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q := NewQueueImp([]int{1}, 10)

		q.EnQueue(2)
		assert.Equal(t, "[1, 2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2}, 10)

		q.EnQueue(3)
		assert.Equal(t, "[1, 2, 3]", q.String())
	})
}

func TestQueueImp_DeQueue(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)

		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1]", func(t *testing.T) {
		q := NewQueueImp([]int{1}, 10)

		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2}, 10)

		q.DeQueue()
		assert.Equal(t, "[2]", q.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2, 3}, 10)

		q.DeQueue()
		assert.Equal(t, "[2, 3]", q.String())
	})
	t.Run("3 dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2, 3}, 10)

		q.DeQueue()
		q.DeQueue()
		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
	t.Run("4 dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2, 3}, 10)

		q.DeQueue()
		q.DeQueue()
		q.DeQueue()
		q.DeQueue()
		assert.Equal(t, "[]", q.String())
	})
}

func TestQueueImp_Front(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)
		assert.Equal(t, -1, q.Front())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2, 3}, 10)
		assert.Equal(t, 1, q.Front())
	})
	t.Run("enqueue [1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)

		assert.Equal(t, 1, q.Front())
	})
	t.Run("enqueue/dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)
		q.DeQueue()
		q.DeQueue()
		assert.Equal(t, 3, q.Front())
	})
}

func TestQueueImp_Rear(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)
		assert.Equal(t, -1, q.Rear())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{1, 2, 3}, 10)
		assert.Equal(t, 3, q.Rear())
	})
	t.Run("enqueue [1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)

		assert.Equal(t, 3, q.Rear())
	})
	t.Run("enqueue/dequeue [1, 2, 3]", func(t *testing.T) {
		q := NewQueueImp([]int{}, 10)
		q.EnQueue(1)
		q.EnQueue(2)
		q.EnQueue(3)
		q.DeQueue()
		q.DeQueue()

		assert.Equal(t, 3, q.Rear())
	})
}

func TestQueueImp(t *testing.T) {
	q := NewQueueImp([]int{}, 5)

	q.EnQueue(1)
	q.EnQueue(2)
	q.EnQueue(3)
	q.EnQueue(4)
	q.EnQueue(5)

	q.DeQueue()
	q.DeQueue()
	q.DeQueue()
	q.DeQueue()

	q.EnQueue(6)
	q.EnQueue(7)
	q.EnQueue(8)
	q.EnQueue(9)
	q.EnQueue(10)

	assert.Equal(t, "[5, 6, 7, 8, 9, 10]", q.String())
}

func BenchmarkQueue(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueue([]int{})

		for j := 0; j < 15000; j++ {
			q.EnQueue(1)
		}
		for j := 0; j < 5000; j++ {
			q.DeQueue()
		}
		for j := 0; j < 15000; j++ {
			q.EnQueue(1)
		}
		for j := 0; j < 15000; j++ {
			q.DeQueue()
		}
		for j := 0; j < 15000; j++ {
			q.EnQueue(1)
		}
	}
}

func BenchmarkQueueImp(b *testing.B) {
	for i := 0; i < b.N; i++ {
		q := NewQueueImp([]int{}, 1000)

		for j := 0; j < 15000; j++ {
			q.EnQueue(1)
		}
		for j := 0; j < 5000; j++ {
			q.DeQueue()
		}
		for j := 0; j < 15000; j++ {
			q.EnQueue(1)
		}
		for j := 0; j < 15000; j++ {
			q.DeQueue()
		}
		for j := 0; j < 15000; j++ {
			q.EnQueue(1)
		}
	}
}
