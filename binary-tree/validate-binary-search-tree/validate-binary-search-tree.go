/*
Given a binary tree, determine if it is a valid binary search tree (BST).

Assume a BST is defined as follows:

The left subtree of a node contains only nodes with keys less than the node's key.
The right subtree of a node contains only nodes with keys greater than the node's key.
Both the left and right subtrees must also be binary search trees.


Example 1:

    2
   / \
  1   3

Input: [2,1,3]
Output: true
Example 2:

    5
   / \
  1   4
     / \
    3   6

Input: [5,1,4,null,null,3,6]
Output: false
Explanation: The root node's value is 5 but its right child's value is 4.
*/

package validate_binary_search_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Recursion
func isValidBST(n *TreeNode) bool {
	return checker(n, nil, nil)
}

/*
Left ->  lower=lower;upper=val
Right -> lower=val;upper=upper
*/
func checker(n *TreeNode, lower *int, upper *int) bool {
	if n == nil {
		return true
	}
	val := n.Val
	if lower != nil && val <= *lower {
		return false
	}
	if upper != nil && val >= *upper {
		return false
	}
	return checker(n.Left, lower, &val) && checker(n.Right, &val, upper)
}

// Iteration
func isValidBST_iter(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := &stack{}
	lower := &stack{}
	upper := &stack{}
	st.push(root)
	lower.push(nil)
	upper.push(nil)

	for !st.isEmpty() {
		n := st.pop()
		l := lower.pop()
		u := upper.pop()

		if l != nil && n.Val <= l.Val {
			return false
		}
		if u != nil && n.Val >= u.Val {
			return false
		}

		if n.Right != nil {
			st.push(n.Right)
			lower.push(n)
			upper.push(u)
		}
		if n.Left != nil {
			st.push(n.Left)
			lower.push(l)
			upper.push(n)
		}
	}

	return true
}

type stack struct {
	data []*TreeNode
}

func (st *stack) push(v *TreeNode) {
	st.data = append(st.data, v)
}
func (st *stack) pop() *TreeNode {
	v := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return v
}
func (st *stack) isEmpty() bool {
	return len(st.data) == 0
}

const MaxUint = ^uint(0)
const MinUint = 0
const MaxInt = int(MaxUint >> 1)
const MinInt = -MaxInt - 1

// InOrder
func isValidBST_inorder(root *TreeNode) bool {
	if root == nil {
		return true
	}
	st := &stack{}
	val := MinInt

	for !st.isEmpty() || root != nil {
		for root != nil {
			st.push(root)
			root = root.Left
		}

		root = st.pop()
		if val >= root.Val {
			return false
		}
		val = root.Val

		root = root.Right
	}
	return true
}
