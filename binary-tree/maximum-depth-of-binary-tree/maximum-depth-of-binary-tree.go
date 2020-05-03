/*
Given a binary tree, find its maximum depth.

The maximum depth is the number of nodes along the longest path from the root node down to the farthest leaf node.

Note: A leaf is a node with no children.

Example:

Given binary tree [3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
return its depth = 3.
*/

package maximum_depth_of_binary_tree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// "Top-Down"

var DEPTH int

func maxDepth_TopDown(root *TreeNode) int {
	DEPTH = 0
	maxDepthTopDown(root, DEPTH+1)
	return DEPTH
}

func maxDepthTopDown(node *TreeNode, depth int) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if DEPTH < depth {
			DEPTH = depth
		}
	}
	maxDepthTopDown(node.Left, depth+1)
	maxDepthTopDown(node.Right, depth+1)
}

// "Bottom-Up"
