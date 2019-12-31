package doubly_circular

import (
	"bytes"
	"errors"
	"strconv"
)

var (
	ErrEmptyList       = errors.New("list is empty")
	ErrIndexOutOfRange = errors.New("index out of range")
	ErrNodeNotFound    = errors.New("node not found")
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

func (node *Node) Equals(other *Node) bool {
	return node == other || node.value == other.value
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
		if curr != l.head {
			b.WriteString(", ")
		} else {
			break
		}
	}
	b.WriteString("]")
	return b.String()
}

func (l *List) ReverseString() string {
	b := bytes.Buffer{}
	curr := l.tail

	b.WriteString("[")
	for curr != nil {
		b.WriteString(strconv.Itoa(curr.value))

		curr = curr.prev
		if curr != l.tail {
			b.WriteString(", ")
		} else {
			break
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

	if l.head == nil { // list is empty
		node.prev, node.next = node, node
		l.head, l.tail = node, node
	} else {
		node.prev, node.next = l.head.prev, l.head
		l.head.prev = node
		l.tail.next = node
		l.head = node
	}
	l.count++
}
