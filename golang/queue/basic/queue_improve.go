package queue

import (
	"strconv"
	"strings"
)

type queueImp struct {
	q       []int
	head    int
	tail    int
	maxHead int
}

func NewQueueImp(val []int, maxHead int) *queueImp {
	q := queueImp{
		q:       val[:],
		head:    -1,
		tail:    len(val) - 1,
		maxHead: maxHead,
	}
	if len(val) > 0 {
		q.head = 0
	}
	return &q
}

func (q *queueImp) EnQueue(val int) {
	q.q = append(q.q, val)
	q.tail++
	if q.head == -1 {
		q.head = 0
	}
}

func (q *queueImp) DeQueue() {
	if q.head == q.tail {
		q.q = q.q[:0]
		q.tail, q.head = -1, -1
		return
	}

	q.head++

	if q.head >= q.maxHead {
		copy(q.q[:], q.q[q.head:q.tail+1])
		q.tail = q.tail - q.head
		q.head = 0
		q.q = q.q[:q.tail+1]
	}
}

func (q *queueImp) Front() int {
	if q.head == -1 {
		return -1
	}
	return q.q[q.head]
}

func (q *queueImp) Rear() int {
	if q.head == -1 {
		return -1
	}
	return q.q[q.tail]
}

const emptySlice = "[]"

func (q *queueImp) String() string {
	if q.head == -1 {
		return emptySlice
	}

	var buf strings.Builder
	buf.WriteString("[")

	for i, v := range q.q[q.head : q.tail+1] {
		buf.WriteString(strconv.Itoa(v))
		if i != q.tail-q.head {
			buf.WriteString(", ")
		}
	}
	//for i := q.head; i <= q.tail; i++ {
	//	buf.WriteString(strconv.Itoa(q.q[i]))
	//	if i != len(q.q)-1 {
	//		buf.WriteString(", ")
	//	}
	//}

	buf.WriteString("]")
	return buf.String()
}
