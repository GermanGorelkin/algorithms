/*
Given a binary tree, return the inorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,3,2]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package binary_tree_inorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	s = append(s, inorderTraversal(root.Left)...)
	s = append(s, root.Val)
	s = append(s, inorderTraversal(root.Right)...)

	return s
}

func inorderTraversal_iteratively(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	st := &stack{}
	cur := root

	for cur != nil || !st.isEmpty() {
		if cur != nil {
			st.push(cur)
			cur = cur.Left
		} else {
			cur = st.pop()
			s = append(s, cur.Val)
			cur = cur.Right
		}
	}

	return s
}

func inorderTraversal_iteratively2(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	st := &stack{}
	cur := root

	for cur != nil || !st.isEmpty() {
		for cur != nil {
			st.push(cur)
			cur = cur.Left
		}

		cur = st.pop()
		s = append(s, cur.Val)
		cur = cur.Right
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
