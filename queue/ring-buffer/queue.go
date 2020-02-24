package ring_buffer

import (
	"strconv"
	"strings"
)

type queue struct {
	data []int
	head int
	tail int
	size int
}

func NewQueue(size int) *queue {
	return &queue{
		data: make([]int, size),
		head: -1,
		tail: -1,
		size: size,
	}
}

func (q *queue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.head == q.tail {
		q.head, q.tail = -1, -1
		return true
	}

	q.head = q.nextIndex(q.head)
	return true
}

func (q *queue) EnQueue(val int) bool {
	if q.IsFull() {
		return false
	}

	q.tail = q.nextIndex(q.tail)
	q.data[q.tail] = val

	if q.IsEmpty() {
		q.head = 0
	}

	return true
}

func (q *queue) nextIndex(ind int) int {
	return (ind + 1) % q.size
}

func (q *queue) IsEmpty() bool {
	return q.head == -1
}

func (q *queue) IsFull() bool {
	return q.nextIndex(q.tail) == q.head
}

func (q *queue) Front() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.head]
}

func (q *queue) Rear() int {
	if q.IsEmpty() {
		return -1
	}

	return q.data[q.tail]
}

const emptySlice = "[]"

func (q *queue) String() string {
	if q.IsEmpty() {
		return emptySlice
	}

	var buf strings.Builder
	buf.WriteString("[")

	if q.head <= q.tail {
		for i, v := range q.data[q.head : q.tail+1] {
			buf.WriteString(strconv.Itoa(v))
			if i != q.tail-q.head {
				buf.WriteString(", ")
			}
		}
	} else { // [4, 5, 1, 2, 3] head=2 tail =1
		// 1, 2, 3,
		for _, v := range q.data[q.head:] {
			buf.WriteString(strconv.Itoa(v))
			buf.WriteString(", ")
		}
		// 4, 5
		for i, v := range q.data[:q.tail+1] {
			buf.WriteString(strconv.Itoa(v))
			if i != q.tail {
				buf.WriteString(", ")
			}
		}
	}

	buf.WriteString("]")
	return buf.String()
}
