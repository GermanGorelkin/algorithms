package smallest_string_starting_from_leaf

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_smallestFromLeaf(t *testing.T) {
	root := &TreeNode{
		Val:   0,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}

	assert.Equal(t, "ba", smallestFromLeaf(root))
}
