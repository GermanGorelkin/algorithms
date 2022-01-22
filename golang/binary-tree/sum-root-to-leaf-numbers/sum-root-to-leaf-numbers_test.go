package sum_root_to_leaf_numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_sumNumbers(t *testing.T) {
	t.Run("[1,2,3]", func(t *testing.T) {
		root := &TreeNode{
			Val:   1,
			Left:  &TreeNode{Val: 2},
			Right: &TreeNode{Val: 3},
		}

		assert.Equal(t, 25, sumNumbers(root))
	})
}
