package queue

import (
	"strconv"
	"strings"
)

type queue struct {
	q []int
}

func NewQueue(val []int) *queue {
	return &queue{val[:]}
}
func (q *queue) EnQueue(val int) {
	q.q = append(q.q, val)
}
func (q *queue) DeQueue() {
	if len(q.q) == 0 {
		return
	}
	q.q = q.q[1:]
}
func (q *queue) Front() int {
	if len(q.q) == 0 {
		return -1
	}
	return q.q[0]
}
func (q *queue) Rear() int {
	if len(q.q) == 0 {
		return -1
	}
	return q.q[len(q.q)-1]
}
func (q *queue) String() string {
	var buf strings.Builder
	buf.WriteString("[")

	for i, v := range q.q {
		buf.WriteString(strconv.Itoa(v))
		if i != len(q.q)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString("]")
	return buf.String()
}
