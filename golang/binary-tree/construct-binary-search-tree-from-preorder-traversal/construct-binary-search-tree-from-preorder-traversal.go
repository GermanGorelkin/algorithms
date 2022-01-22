/*
Return the root node of a binary search tree that matches the given preorder traversal.

(Recall that a binary search tree is a binary tree where for every node, any descendant of node.left has a value < node.val, and any descendant of node.right has a value > node.val.  Also recall that a preorder traversal displays the value of the node first, then traverses node.left, then traverses node.right.)

It's guaranteed that for the given test cases there is always possible to find a binary search tree with the given requirements.

Example 1:

Input: [8,5,1,7,10,12]
Output: [8,5,10,1,7,null,12]


Constraints:

1 <= preorder.length <= 100
1 <= preorder[i] <= 10^8
The values of preorder are distinct.
 */

package construct_binary_search_tree_from_preorder_traversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func bstFromPreorder(preorder []int) *TreeNode {
	root := &TreeNode{}
	root.Val = preorder[0]
	st := &stack{}
	st.push(root)

	for i := 1; i < len(preorder); i++ {
		node := &TreeNode{Val: preorder[i]}
		prev := st.pop()
		if preorder[i] < preorder[i-1] {
			prev.Left = node
			st.push(prev)
			st.push(node)
		} else {
			prev.Right = node
		}
	}

	return root
}

type stack struct {
	data []*TreeNode
}
func (s *stack) push(v *TreeNode) {
	s.data = append(s.data, v)
}
func (s *stack) pop() *TreeNode {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}