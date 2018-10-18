package singly

import (
	"bytes"
	"strconv"
)

type Node struct {
	value int
	next  *Node
}

type List struct {
	head  *Node
	count int
}

func NewList() *List {
	return &List{}
}

func (l *List) Size() int {
	return l.count
}

func (l *List) IsEmpty() bool {
	return l.count == 0
}

func (l *List) AddHead(value int) {
	l.head = &Node{next: l.head, value: value}
	l.count++
}

func (l *List) AddTail(value int) {
	n := &Node{next: nil, value: value}
	c := l.head

	if c == nil {
		l.head = n
		l.count++
		return
	}

	for c.next != nil {
		c = c.next
	}

	c.next = n
	l.count++
}

func (l *List) String() string {
	b := bytes.Buffer{}
	c := l.head

	b.WriteString("[")
	for c != nil {
		b.WriteString(strconv.Itoa(c.value))
		c = c.next
		if c != nil {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}

func (l *List) SortedInsert(value int) {
	n := &Node{next: nil, value: value}
	c := l.head

	if c == nil || c.value > value {
		n.next = c
		l.head = n
		return
	}

	for c.next != nil && c.next.value < value {
		c = c.next
	}

	n.next = c.next
	c.next = n
}
