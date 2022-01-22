package n_ary_tree_postorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	for _, node := range root.Children {
		res = append(res, postorder(node)...)
	}

	return append(res, root.Val)
}
