package implement_trie_prefix_tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Insert(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		trie := Constructor()

		trie.Insert("a")
		assert.Equal(t, []byte{'a'}, TrieTraversal(trie.root))

		trie.Insert("b")
		assert.Equal(t, []byte{'a', 'b'}, TrieTraversal(trie.root))

		trie.Insert("bb")
		assert.Equal(t, []byte{'a', 'b', 'b'}, TrieTraversal(trie.root))

		trie.Insert("apple")
		assert.Equal(t, []byte{'a', 'p', 'p', 'l', 'e', 'b', 'b'}, TrieTraversal(trie.root))

		trie.Insert("bar")
		assert.Equal(t, []byte{'a', 'p', 'p', 'l', 'e', 'b', 'a', 'r', 'b'}, TrieTraversal(trie.root))
	})
}

func TestTrieTraversal(t *testing.T) {
	t.Run("1", func(t *testing.T) {

		n1 := &node{char: 'a'}
		n2 := &node{char: 'b'}
		n3 := &node{char: 'c'}
		n4 := &node{char: 'd'}
		n5 := &node{char: 'e'}
		n6 := &node{char: 'f'}

		n1.children[1] = n2
		n1.children[2] = n3
		n1.children[3] = n4

		n2.children[4] = n5
		n2.children[5] = n6

		root := &node{}
		root.children[0] = n1

		want := []byte{'a', 'b', 'e', 'f', 'c', 'd'}

		assert.Equal(t, want, TrieTraversal(root))
	})
}
