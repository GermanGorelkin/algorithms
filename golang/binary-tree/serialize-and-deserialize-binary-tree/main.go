package main

import (
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct{}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this Codec) serialize(root *TreeNode) string {
	if root == nil {
		return "[]"
	}

	ss := []string{}

	q := new(queue)
	q.enQueue(root)

	for !q.isEmpty() {
		node := q.deQueue()
		if node == nil {
			ss = append(ss, "null")
		} else {
			q.enQueue(node.Left)
			q.enQueue(node.Right)

			ss = append(ss, strconv.Itoa(node.Val))
		}
	}

	for i := len(ss) - 1; i >= 0 && ss[i] == "null"; i-- {
		ss = ss[:len(ss)-1]
	}

	result := "[" + strings.Join(ss, ",") + "]"

	return result
}

// Deserializes your encoded data to tree.
/*
	left  = 2i+1
	right = 2i+2
*/
func (this Codec) deserialize(data string) *TreeNode {
	if data == "[]" {
		return nil
	}
	data = data[1 : len(data)-1] // remove []
	ss := strings.Split(data, ",")

	q := new(queue)
	root := &TreeNode{Val: mustConvertToInt(ss[0])}
	q.enQueue(root)
	index := 0

	for !q.isEmpty() {
		node := q.deQueue()
		if 2*index+1 < len(ss) && ss[2*index+1] != "null" {
			lnode := &TreeNode{Val: mustConvertToInt(ss[2*index+1])}
			node.Left = lnode
			q.enQueue(lnode)

		}
		if 2*index+2 < len(ss) && ss[2*index+2] != "null" {
			rnode := &TreeNode{Val: mustConvertToInt(ss[2*index+2])}
			node.Right = rnode
			q.enQueue(rnode)
		}
		index++
	}

	return root
}

func mustConvertToInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

//-----queue

type queue struct {
	data []*TreeNode
}

func (q *queue) enQueue(v *TreeNode) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() *TreeNode {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}
