package leafsimilartrees

import "reflect"

/*
	№872
	https://leetcode.com/problems/leaf-similar-trees/
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	leafs1 := make([]int, 1)
	leafs2 := make([]int, 1)

	getLeafs(root1, &leafs1)
	getLeafs(root2, &leafs2)

	return reflect.DeepEqual(leafs1, leafs2)
}

func getLeafs(root *TreeNode, leafs *[]int) {
	if root == nil {
		return
	}

	if root.Left == nil && root.Right == nil {
		*leafs = append(*leafs, root.Val)
		return
	}

	getLeafs(root.Left, leafs)
	getLeafs(root.Right, leafs)
}
