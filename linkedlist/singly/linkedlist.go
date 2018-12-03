package singly

import (
	"bytes"
	"errors"
	"strconv"
)

var ErrEmptyList = errors.New("list is empty")

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

func (l *List) IsPresent(data int) bool {
	t := l.head
	for t != nil {
		if t.value == data {
			return true
		}
		t = t.next
	}
	return false
}

func (l *List) RemoveHead() (int, error) {
	if l.IsEmpty() {
		return 0, ErrEmptyList
	}
	value := l.head.value
	l.head = l.head.next
	l.count--
	return value, nil
}

// DeleteNode удаляет первый найденнный элемент который равен delValue
func (l *List) DeleteNode(delValue int) (bool, error) {
	if l.IsEmpty() {
		return false, ErrEmptyList
	}

	if l.head.value == delValue {
		l.head = l.head.next
		l.count--
		return true, nil
	}

	t := l.head
	for t.next != nil {
		if t.next.value == delValue {
			t.next = t.next.next
			l.count--
			return true, nil
		}
		t = t.next
	}

	return false, nil
}

// DeleteNode удаляет все элементы которые равены delValue
func (l *List) DeleteNodes(delValue int) (isDel bool, err error) {
	if l.IsEmpty() {
		return false, ErrEmptyList
	}

	if l.head.value == delValue {
		l.head = l.head.next
		l.count--
		isDel = true
	}

	t := l.head
	for t != nil && t.next != nil {
		if t.next.value == delValue {
			t.next = t.next.next
			l.count--
			isDel = true
		}
		t = t.next
	}

	return
}

func (l *List) FreeList() {
	l.count = 0
	l.head = nil
}

func (l *List) Reverse() {
	curr := l.head
	var prev, next *Node

	for curr != nil {
		next = curr.next
		curr.next = prev
		prev = curr
		curr = next
	}

	l.head = prev
}
