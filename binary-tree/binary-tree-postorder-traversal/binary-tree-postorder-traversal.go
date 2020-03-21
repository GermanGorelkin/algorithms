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
