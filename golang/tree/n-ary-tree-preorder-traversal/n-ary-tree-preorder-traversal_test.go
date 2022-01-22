package n_ary_tree_preorder_traversal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_preorder(t *testing.T) {
	t.Run("1", func(t *testing.T) {

		n1 := &Node{Val: 1}
		n2 := &Node{Val: 2}
		n3 := &Node{Val: 3}
		n4 := &Node{Val: 4}
		n5 := &Node{Val: 5}
		n6 := &Node{Val: 6}

		n1.Children = append(n1.Children, n3, n2, n4)
		n3.Children = append(n3.Children, n5, n6)

		want := []int{1, 3, 5, 6, 2, 4}

		assert.Equal(t, want, preorder(n1))
	})

	t.Run("1", func(t *testing.T) {

		n1 := &Node{Val: 1}
		n2 := &Node{Val: 2}
		n3 := &Node{Val: 3}
		n4 := &Node{Val: 4}
		n5 := &Node{Val: 5}
		n6 := &Node{Val: 6}
		n7 := &Node{Val: 7}
		n8 := &Node{Val: 8}
		n9 := &Node{Val: 9}
		n10 := &Node{Val: 10}
		n11 := &Node{Val: 11}
		n12 := &Node{Val: 12}
		n13 := &Node{Val: 13}
		n14 := &Node{Val: 14}

		n1.Children = append(n1.Children, n2, n3, n4, n5)
		n3.Children = append(n3.Children, n6, n7)
		n7.Children = append(n7.Children, n11)
		n11.Children = append(n11.Children, n14)
		n4.Children = append(n4.Children, n8)
		n8.Children = append(n8.Children, n12)
		n5.Children = append(n5.Children, n9, n10)
		n9.Children = append(n9.Children, n13)

		want := []int{1, 2, 3, 6, 7, 11, 14, 4, 8, 12, 5, 9, 13, 10}

		assert.Equal(t, want, preorder(n1))
	})
}
