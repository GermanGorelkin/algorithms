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

func (l *List) AddTail(val int) {
	node := &Node{value: val}

	if l.tail == nil {
		node.prev, node.next = node, node
		l.head, l.tail = node, node
	} else {
		node.prev, node.next = l.tail, l.tail.next
		l.head.prev = node
		l.tail.next = node
		l.tail = node
	}
	l.count++
}

func (l *List) RemoveHead() (*Node, error) {
	if l.IsEmpty() {
		return nil, ErrEmptyList
	}

	head := l.head
	if l.head == l.head.next { // 1 node
		l.head, l.tail = nil, nil
	} else {
		l.tail.next = l.head.next
		l.head.next.prev = l.tail
		l.head = l.head.next
	}
	l.count--

	return head, nil
}

func (l *List) RemoveTail() (*Node, error) {
	if l.IsEmpty() {
		return nil, ErrEmptyList
	}

	tail := l.tail
	if l.tail == l.tail.next { // 1 node
		l.head, l.tail = nil, nil
	} else {
		l.head.prev = l.tail.prev
		l.tail.prev.next = l.head
		l.tail = l.tail.prev
	}
	l.count--

	return tail, nil
}
