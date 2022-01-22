package v2

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	depthX, depthY   int
	parentX, parentY *TreeNode
)

func isCousins(root *TreeNode, x int, y int) bool {
	dfs(root, nil, x, y, 0)
	return depthX == depthY && parentX != parentY
}

func dfs(node, parent *TreeNode, x, y, depth int) {
	if node == nil {
		return
	}
	if node.Val == x {
		depthX = depth
		parentX = parent
	} else if node.Val == y {
		depthY = depth
		parentY = parent
	}
	dfs(node.Left, node, x, y, depth+1)
	dfs(node.Right, node, x, y, depth+1)
}
