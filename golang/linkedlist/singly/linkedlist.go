package singly

import (
	"bytes"
	"errors"
	"strconv"
)

var ErrEmptyList = errors.New("list is empty")
var ErrIndexOutOfRange = errors.New("index out of range")

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

func (l *List) AddNodeAtTail(node Node) {
	c := l.head

	if c == nil {
		l.head = &node
		l.count++
		return
	}

	for c.next != nil {
		c = c.next
	}

	c.next = &node
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

func (l *List) ReverseRecurse() {
	l.head = l.reverseRecurseUtil(l.head, nil)
}
func (l *List) reverseRecurseUtil(curr *Node, next *Node) *Node {
	var ret *Node
	if curr == nil {
		return nil
	}
	if curr.next == nil {
		curr.next = next
		return curr
	}
	ret = l.reverseRecurseUtil(curr.next, curr)
	curr.next = next
	return ret
}

func (l *List) RemoveDuplicate() {
	curr := l.head
	for curr != nil {
		if curr.next != nil && curr.value == curr.next.value {
			curr.next = curr.next.next
		} else {
			curr = curr.next
		}
	}
}

func (l *List) CopyListReversed() *List {
	var temp, temp2 *Node
	curr := l.head
	for curr != nil {
		temp2 = &Node{curr.value, temp}
		curr = curr.next
		temp = temp2
	}
	ll2 := new(List)
	ll2.head = temp
	return ll2
}

func (l *List) CopyList() *List {
	var head, tail, temp *Node
	curr := l.head
	if curr == nil {
		ll2 := new(List)
		ll2.head = nil
		return ll2
	}
	head = &Node{curr.value, nil}
	tail = head
	curr = curr.next
	for curr != nil {
		temp = &Node{curr.value, nil}
		tail.next = temp
		tail = temp
		curr = curr.next
	}
	ll2 := new(List)
	ll2.head = head
	return ll2
}

func (l *List) CompareList(ll *List) bool {
	return l.compareListUtil(l.head, ll.head)
}
func (l *List) compareListUtil(head1 *Node, head2 *Node) bool {
	if head1 == nil && head2 == nil {
		return true
	} else if (head1 == nil) || (head2 == nil) || (head1.value != head2.value) {
		return false
	} else {
		return l.compareListUtil(head1.next, head2.next)
	}
}

func (l *List) FindLength() int {
	curr := l.head
	count := 0
	for curr != nil {
		count++
		curr = curr.next
	}
	return count
}

func (l *List) NthNodeFromBeginning(index int) (int, error) {
	if index > l.Size() || index < 1 {
		return 0, ErrIndexOutOfRange
	}

	count := 0
	curr := l.head
	for curr != nil && count < index-1 {
		count++
		curr = curr.next
	}
	return curr.value, nil
}

func (l *List) NthNodeFromEnd(index int) (int, error) {
	size := l.Size()
	if size != 0 && size < index {
		return 0, ErrIndexOutOfRange
	}

	startIndex := size - index + 1
	return l.NthNodeFromBeginning(startIndex)
}

func (l *List) LoopDetect() bool {
	slowPtr := l.head
	fastPtr := l.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			return true
		}
	}
	return false
}

func (l *List) ReverseListLoopDetect() bool {
	tempHead := l.head
	l.Reverse()
	if tempHead == l.head {
		l.Reverse()
		return true
	}
	l.Reverse()
	return false
}

func (l *List) LoopTypeDetect() int {
	slowPtr := l.head
	fastPtr := l.head
	for fastPtr.next != nil && fastPtr.next.next != nil {
		if l.head == fastPtr.next || l.head == fastPtr.next.next {
			return 2
		}
		slowPtr = slowPtr.next
		fastPtr = fastPtr.next.next
		if slowPtr == fastPtr {
			return 1
		}
	}
	return 0
}

func (l *List) LoopPointDetect() *Node {
	slow := l.head
	fast := l.head
	for fast.next != nil && fast.next.next != nil {
		slow = slow.next
		fast = fast.next.next

		if slow == fast {
			if fast == l.head { // CircularLoop
				for slow.next != fast {
					slow = slow.next
				}
			} else {
				fast = l.head
				for fast.next != slow.next {
					fast = fast.next
					slow = slow.next
				}
			}
			return slow
		}
	}
	return nil
}

func (l *List) RemoveLoop() {
	loopPoint := l.LoopPointDetect()
	if loopPoint == nil {
		return
	}

	loopPoint.next = nil
}

func (l *List) FindIntersection(head2 *Node) *Node {
	head := l.head
	l1 := 0
	l2 := 0
	tempHead := head
	tempHead2 := head2
	for tempHead != nil {
		l1++
		tempHead = tempHead.next
	}
	for tempHead2 != nil {
		l2++
		tempHead2 = tempHead2.next
	}
	var diff int
	if l1 < 12 {
		temp := head
		head = head2
		head2 = temp
		diff = l2 - l1
	} else {
		diff = l1 - l2
	}
	for ; diff > 0; diff-- {
		head = head.next
	}
	for head != head2 {
		head = head.next
		head2 = head2.next
	}
	return head
}

func (l *List) OddEvenList() {
	odd, even, evenHead := l.head, l.head.next, l.head.next

	for even != nil && even.next != nil {
		odd.next = even.next
		odd = odd.next
		even.next = odd.next
		even = even.next
	}
	odd.next = evenHead
}

func (l *List) IsPalindrome() bool {
	if l.head == nil || l.head.next == nil {
		return false
	}

	slow, fast := l.head, l.head.next.next

	for fast != nil && fast.next != nil {
		slow = slow.next
		fast = fast.next.next
	}

	if fast != nil {
		return false
		//slow = slow.next
	}

	var prev, next *Node
	for slow != nil {
		next = slow.next
		slow.next = prev
		prev = slow
		slow = next
	}
	slow = prev

	fast = l.head
	for slow != nil && fast != nil {
		if slow.value != fast.value {
			return false
		}
		slow = slow.next
		fast = fast.next
	}

	return true
}

func (l *List) AddTwoNumbers(l2 *Node) *Node {
	head := &Node{}
	cur := head
	carry := 0

	l1 := l.head

	for l1 != nil || l2 != nil {
		if l1 != nil {
			carry += l1.value
			l1 = l1.next
		}
		if l2 != nil {
			carry += l2.value
			l2 = l2.next
		}

		val := carry % 10
		carry = carry / 10

		cur.next = &Node{value: val}
		cur = cur.next
	}

	if carry > 0 {
		cur.next = &Node{value: carry}
	}

	return head.next
}

func (l *List) RotateRight(k int) {
	if l.IsEmpty() || k == 0 {
		return
	}

	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = l.head

	findIndex := l.count - k%l.count
	for findIndex > 0 {
		curr = curr.next
		findIndex--
	}

	l.head = curr.next
	curr.next = nil
}
