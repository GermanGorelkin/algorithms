package sum_root_to_leaf_numbers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumNumbers(root *TreeNode) int {
	return traversal(root, 0)
}

func traversal(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}

	prevSum = prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return prevSum
	}

	return traversal(root.Left, prevSum) + traversal(root.Right, prevSum)
}
