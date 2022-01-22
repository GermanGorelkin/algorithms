package smallest_string_starting_from_leaf

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func smallestFromLeaf(root *TreeNode) string {
	return search(root, "")
}

func search(root *TreeNode, prefix string) string {
	if root == nil {
		return ""
	}

	prefix += string(byte(root.Val) + 'a')

	if root.Left == nil && root.Right == nil {
		return reverse(prefix)
	}

	return min(search(root.Left, prefix), search(root.Right, prefix))
}

func min(a, b string) string {
	if a == "" {
		return b
	}
	if b == "" {
		return a
	}
	if a < b {
		return a
	}
	return b
}

func reverse(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}
