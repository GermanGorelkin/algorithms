/*
Given an n-ary tree, return the preorder traversal of its nodes' values.
*/

package n_ary_tree_preorder_traversal

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, root.Val)

	for _, node := range root.Children {
		res = append(res, preorder(node)...)
	}

	return res
}
