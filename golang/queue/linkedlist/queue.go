package linkedlist

import (
	"strconv"
	"strings"
)

type node struct {
	data int
	next *node
}

type queue struct {
	head *node
	tail *node
	size int
}

func NewQueue() *queue {
	return &queue{}
}

func (q *queue) EnQueue(val int) bool {
	nd := &node{data: val}
	q.size++

	if q.head == nil {
		q.head = nd
	}
	if q.tail != nil {
		q.tail.next = nd
	}
	q.tail = nd

	return true
}

func (q *queue) DeQueue() bool {
	if q.IsEmpty() {
		return false
	}

	if q.head == q.tail {
		q.head, q.tail = nil, nil
	} else {
		q.head = q.head.next
	}
	q.size--

	return true
}

func (q *queue) IsEmpty() bool {
	return q.size == 0
}

func (q *queue) Size() int {
	return q.size
}

func (q *queue) String() string {
	var buf strings.Builder
	buf.WriteString("[")

	cur := q.head
	for cur != nil {
		buf.WriteString(strconv.Itoa(cur.data))

		cur = cur.next
		if cur != nil {
			buf.WriteString(", ")
		}
	}

	buf.WriteString("]")
	return buf.String()
}
