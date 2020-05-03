/*
Given a binary tree, check whether it is a mirror of itself (ie, symmetric around its center).

For example, this binary tree [1,2,2,3,4,4,3] is symmetric:

    1
   / \
  2   2
 / \ / \
3  4 4  3


But the following [1,2,2,null,3,null,3] is not:

    1
   / \
  2   2
   \   \
   3    3


Follow up: Solve it both recursively and iteratively.
*/

package symmetric_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSymmetric_DFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return dfs(root.Left, root.Right)
}

func dfs(l *TreeNode, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return dfs(l.Left, r.Right) && dfs(l.Right, r.Left)
}

func isSymmetric_BFS(root *TreeNode) bool {
	if root == nil {
		return true
	}
	q := &queue{}
	q.enQueue(root)
	q.enQueue(root)

	for !q.isEmpty() {
		n1 := q.deQueue()
		n2 := q.deQueue()

		if n1 == nil && n2 == nil {
			continue
		}
		if n1 == nil || n2 == nil {
			return false
		}
		if n1.Val != n2.Val {
			return false
		}

		q.enQueue(n1.Left)
		q.enQueue(n2.Right)
		q.enQueue(n1.Right)
		q.enQueue(n2.Left)
	}

	return true
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
