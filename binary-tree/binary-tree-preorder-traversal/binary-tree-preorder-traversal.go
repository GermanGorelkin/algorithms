/*
Given a binary tree, return the preorder traversal of its nodes' values.

Example:

Input: [1,null,2,3]
   1
    \
     2
    /
   3

Output: [1,2,3]
Follow up: Recursive solution is trivial, could you do it iteratively?
*/

package binary_tree_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func preorderTraversal(root *TreeNode) (s []int) {
	if root == nil {
		return s
	}

	s = append(s, root.Val)
	s = append(s, preorderTraversal(root.Left)...)
	s = append(s, preorderTraversal(root.Right)...)

	return s
}
