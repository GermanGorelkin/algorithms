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
