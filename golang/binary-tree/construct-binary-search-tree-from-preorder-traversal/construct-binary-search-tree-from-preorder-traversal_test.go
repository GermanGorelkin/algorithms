package construct_binary_search_tree_from_preorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bstFromPreorder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		input := []int{8, 5, 1, 7, 10, 12}

		root := bstFromPreorder(input)

		got := preorder(root)

		assert.Equal(t, input, got)
	})
}

func preorder(root *TreeNode) []int {
	var res []int
	if root == nil {
		return res
	}

	res = append(res, root.Val)
	res = append(res, preorder(root.Left)...)
	res = append(res, preorder(root.Right)...)

	return res
}