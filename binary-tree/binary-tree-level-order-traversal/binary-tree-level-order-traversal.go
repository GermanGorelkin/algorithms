/*
Given a binary tree, return the level order traversal of its nodes' values. (ie, from left to right, level by level).

For example:
Given binary tree [3,9,20,null,null,15,7],
    3
   / \
  9  20
    /  \
   15   7
return its level order traversal as:
[
  [3],
  [9,20],
  [15,7]
]
*/

package binary_tree_level_order_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	res := make([][]int, 0)
	if root == nil {
		return res
	}
	q := &queue{}
	q.EnQueue(root)
	res = append(res, []int{root.Val})

	for !q.IsEmpty() {
		size := q.Size()
		var s []int
		for ; size > 0; size-- {
			p := q.DeQueue()

			if p.Left != nil {
				q.EnQueue(p.Left)
				s = append(s, p.Left.Val)
			}
			if p.Right != nil {
				q.EnQueue(p.Right)
				s = append(s, p.Right.Val)
			}
		}
		if s != nil {
			res = append(res, s)
		}
	}

	return res
}

type queue struct {
	data []*TreeNode
}

func (q *queue) EnQueue(v *TreeNode) {
	q.data = append(q.data, v)
}
func (q *queue) DeQueue() *TreeNode {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) IsEmpty() bool {
	return len(q.data) == 0
}
func (q *queue) Size() int {
	return len(q.data)
}
