package doubly

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
		if curr != nil {
			b.WriteString(", ")
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
		node.prev = l.tail
		l.tail = node
	}

	l.count++
}

func (l *List) AddAtIndex(val, index int) error {
	if l.Size() < index {
		return ErrIndexOutOfRange
	}
	if index < 0 {
		index = 0
	}

	node := &Node{value: val}

	// add at head
	if index == 0 {
		l.AddHead(val)
		return nil
	}
	// add at tail
	if l.Size() == index {
		l.AddTail(val)
		return nil
	}

	pred := l.head
	for i := 0; i < index; i++ {
		pred = pred.next
	}

	node.prev = pred.prev
	pred.prev.next = node
	pred.prev = node
	node.next = pred
	l.count++

	return nil
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

func (l *List) RemoveTail() (int, error) {
	if l.IsEmpty() {
		return 0, ErrEmptyList
	}

	val := l.tail.value
	l.tail = l.tail.prev

	if l.tail == nil {
		l.head = nil
	} else {
		l.tail.next = nil
	}

	l.count--
	return val, nil
}

func (l *List) RemoveNode(node *Node) bool {
	if l.head == nil {
		return false
	}
	curr := l.head

	for curr != nil {
		if node.Equals(curr) {
			if curr.prev == nil { // curr is head
				_, _ = l.RemoveHead()
			} else if curr.next == nil { // curr is tail
				_, _ = l.RemoveTail()
			} else {
				curr.prev.next = curr.next
				curr.next.prev = curr.prev
				l.count--
			}
			return true
		}
		curr = curr.next
	}

	return false
}

func (l *List) Find(value int) (*Node, error) {
	curr := l.head
	for curr != nil {
		if curr.value == value {
			return curr, nil
		}
		curr = curr.next
	}
	return nil, ErrNodeNotFound
}

func (l *List) ReverseList() {
	curr := l.head
	l.head, l.tail = l.tail, l.head

	for curr != nil {
		curr.next, curr.prev = curr.prev, curr.next
		curr = curr.prev
	}
}
