package linkedlist

import (
	"bytes"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Print(l *ListNode) string {
	b := bytes.Buffer{}
	b.WriteString("[")
	for l != nil {
		b.WriteString(strconv.Itoa(l.Val))
		l = l.Next
		if l != nil {
			b.WriteString(", ")
		}
	}
	b.WriteString("]")
	return b.String()
}
