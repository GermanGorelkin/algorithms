/*
Implement a trie with insert, search, and startsWith methods.

Example:

Trie trie = new Trie();

trie.insert("apple");
trie.search("apple");   // returns true
trie.search("app");     // returns false
trie.startsWith("app"); // returns true
trie.insert("app");
trie.search("app");     // returns true
Note:

You may assume that all inputs are consist of lowercase letters a-z.
All inputs are guaranteed to be non-empty strings.
*/

package implement_trie_prefix_tree

type node struct {
	isWord   bool
	char     byte
	children [26]*node
}

type Trie struct {
	root *node
}

/** Initialize your data structure here. */
func Constructor() Trie {
	return Trie{root: &node{}}
}

// Traversal preorder traversals
func TrieTraversal(root *node) []byte {
	var res []byte
	if root == nil {
		return res
	}

	for i, node := range root.children {
		if node == nil {
			continue
		}
		ch := byte(i + 'a')
		res = append(res, ch)
		res = append(res, TrieTraversal(node)...)
	}

	return res
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	root := this.root

	for i := range word {
		if root.children[word[i]-'a'] == nil {
			root.children[word[i]-'a'] = &node{
				char: word[i],
			}
		}
		root = root.children['a'-i]
	}

	root.isWord = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	return false
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {
	return false
}
