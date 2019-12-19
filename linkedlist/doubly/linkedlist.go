package doubly

import (
	"bytes"
	"errors"
	"strconv"
)

var (
	ErrEmptyList       = errors.New("list is empty")
	ErrIndexOutOfRange = errors.New("index out of range")
)

type List struct {
	head  *Node
	tail  *Node
	count int
}

type Node struct {
	value int
	prev  *Node
	next  *Node
}

func NewList() *List {
	return &List{}
}

func (l *List) String() string {
	b := bytes.Buffer{}
	curr := l.head

	b.WriteString("[")
	for curr != nil {
		b.WriteString(strconv.Itoa(curr.value))
		curr = curr.next
		if curr != nil {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (l *List) Size() int {
	return l.count
}

func (l *List) IsEmpty() bool {
	return l.count == 0
}

func (l *List) AddHead(val int) {
	node := &Node{value: val}

	if l.head == nil {
		l.head = node
		l.tail = node
	} else {
		node.next = l.head
		l.head.prev = node
		l.head = node
	}

	l.count++
}

func (l *List) AddTail(val int) {
	node := &Node{value: val}

	if l.tail == nil {
		l.head = node
		l.tail = node
	} else {
		l.tail.next = node
		l.tail = node
	}

	l.count++
}

func (l *List) RemoveHead() (int, error) {
	if l.IsEmpty() {
		return 0, ErrEmptyList
	}

	val := l.head.value
	l.head = l.head.next

	if l.head == nil {
		l.tail = nil
	} else {
		l.head.prev = nil
	}

	l.count--

	return val, nil
}
