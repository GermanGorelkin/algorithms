/*
Given an n-ary tree, return the level order traversal of its nodes' values.

Nary-Tree input serialization is represented in their level order traversal, each group of children is separated by the null value
*/

package n_ary_tree_level_order_traversal

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	q := &queue{}
	q.enQueue(root)

	for !q.isEmpty() {
		size := q.size()
		s := make([]int, 0, size)
		for ; size > 0; size-- {
			node := q.deQueue()
			s = append(s, node.Val)

			for _, child := range node.Children {
				q.enQueue(child)
			}
		}
		res = append(res, s)
	}

	return res
}

type queue struct {
	data []*Node
}

func (q *queue) enQueue(v *Node) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() *Node {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}
func (q *queue) size() int {
	return len(q.data)
}
