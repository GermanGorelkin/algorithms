/*
https://stepik.org/lesson/45970/step/1?unit=24123
Обход дерева
*/

package main

import (
	"bytes"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_inOrder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `4 1 2
		2 3 4
		5 -1 -1
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 5

		root := treeBuild(n, r)

		buf := new(bytes.Buffer)
		inOrder(root, buf)
		assert.Equal(t, "1 2 3 4 5 ", buf.String())
	})
}

func Test_preOrder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `4 1 2
		2 3 4
		5 -1 -1
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 5

		root := treeBuild(n, r)

		buf := new(bytes.Buffer)
		preOrder(root, buf)
		assert.Equal(t, "4 2 1 3 5 ", buf.String())
	})
}

func Test_postOrder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `4 1 2
		2 3 4
		5 -1 -1
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 5

		root := treeBuild(n, r)

		buf := new(bytes.Buffer)
		postOrder(root, buf)
		assert.Equal(t, "1 3 2 5 4 ", buf.String())
	})
}

func Test_treeBuild(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `4 1 2
		2 3 4
		5 -1 -1
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 5

		root := treeBuild(n, r)

		got := treeSerialize(root)
		assert.Equal(t, "[4,2,5,1,3]", got)
	})
	t.Run("2", func(t *testing.T) {
		s := `0 7 2
		10 -1 -1
		20 -1 6
		30 8 9
		40 3 -1
		50 -1 -1
		60 1 -1
		70 5 4
		80 -1 -1
		90 -1 -1`
		r := strings.NewReader(s)
		n := 10

		root := treeBuild(n, r)

		got := treeSerialize(root)
		assert.Equal(t, "[0,70,20,50,40,null,60,null,null,30,null,10,null,80,90]", got)
	})
}

//----- help funcs
func treeSerialize(root *node) string {
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
			q.enQueue(node.left)
			q.enQueue(node.rigth)

			ss = append(ss, strconv.Itoa(node.key))
		}
	}

	for i := len(ss) - 1; i >= 0 && ss[i] == "null"; i-- {
		ss = ss[:len(ss)-1]
	}

	result := "[" + strings.Join(ss, ",") + "]"

	return result
}

//-----queue
type queue struct {
	data []*node
}

func (q *queue) enQueue(v *node) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() *node {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}
