/*
Given a binary tree, return the postorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [3,2,1]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package binary_tree_postorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func postorderTraversal(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	s = append(s, postorderTraversal(root.Left)...)
	s = append(s, postorderTraversal(root.Right)...)
	s = append(s, root.Val)

	return s
}

func postorderTraversal_iteratively(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	st := &stack{}
	cur := root

	for cur != nil || !st.isEmpty() {
		if cur != nil {
			st.push(cur)
			s = append([]int{cur.Val}, s...)
			cur = cur.Right
		} else {
			cur = st.pop()
			cur = cur.Left
		}
	}

	return s
}

type stack struct {
	s []*TreeNode
}

func (s *stack) push(val *TreeNode) {
	s.s = append(s.s, val)
}
func (s *stack) pop() *TreeNode {
	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return val
}
func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}
