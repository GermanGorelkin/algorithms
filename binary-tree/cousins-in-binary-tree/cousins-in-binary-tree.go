/*
In a binary tree, the root node is at depth 0, and children of each depth k node are at depth k+1.

Two nodes of a binary tree are cousins if they have the same depth, but have different parents.

We are given the root of a binary tree with unique values, and the values x and y of two different nodes in the tree.

Return true if and only if the nodes corresponding to the values x and y are cousins.



Example 1:


Input: root = [1,2,3,4], x = 4, y = 3
Output: false
Example 2:


Input: root = [1,2,3,null,4,null,5], x = 5, y = 4
Output: true
Example 3:



Input: root = [1,2,3,null,4], x = 2, y = 3
Output: false


Note:

The number of nodes in the tree will be between 2 and 100.
Each node has a unique integer value from 1 to 100.

*/

package cousins_in_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isCousins(root *TreeNode, x int, y int) bool {
	q := &queue{}
	q.enQueue(root)

	for !q.isEmpty() {
		size := q.size()
		foundX, foundY := false, false

		for ; size > 0; size-- {
			node := q.deQueue()

			if childEqualTo(node, x) && childEqualTo(node, y) {
				return false
			}

			if node.Val == x {
				foundX = true
			} else if node.Val == y {
				foundY = true
			}
			if foundX && foundY {
				return true
			}

			if node.Left != nil {
				q.enQueue(node.Left)
			}
			if node.Right != nil {
				q.enQueue(node.Right)
			}
		}
	}

	return false
}

func childEqualTo(root *TreeNode, v int) bool {
	return (root.Left != nil && root.Left.Val == v) || (root.Right != nil && root.Right.Val == v)
}

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
func (q *queue) size() int {
	return len(q.data)
}
