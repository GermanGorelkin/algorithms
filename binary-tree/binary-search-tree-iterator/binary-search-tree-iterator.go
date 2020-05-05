/*
Implement an iterator over a binary search tree (BST). Your iterator will be initialized with the root node of a BST.

Calling next() will return the next smallest number in the BST.



Example:
	          7
			/   \
           3     15
			    /  \
               9    20


BSTIterator iterator = new BSTIterator(root);
iterator.next();    // return 3
iterator.next();    // return 7
iterator.hasNext(); // return true
iterator.next();    // return 9
iterator.hasNext(); // return true
iterator.next();    // return 15
iterator.hasNext(); // return true
iterator.next();    // return 20
iterator.hasNext(); // return false


Note:

next() and hasNext() should run in average O(1) time and uses O(h) memory, where h is the height of the tree.
You may assume that next() call will always be valid, that is, there will be at least a next smallest number in the BST when next() is called.
*/

package binary_search_tree_iterator

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	root    *TreeNode
	curNode *TreeNode
	st      *stack
}

func Constructor(root *TreeNode) BSTIterator {
	st := &stack{}
	iter := BSTIterator{
		root:    root,
		curNode: root,
		st:      st,
	}
	return iter
}

// Next returns the next smallest number
func (this *BSTIterator) Next() int {
	var nextVal int
	for this.curNode != nil {
		this.st.push(this.curNode)
		this.curNode = this.curNode.Left
	}

	this.curNode = this.st.pop()
	nextVal = this.curNode.Val
	this.curNode = this.curNode.Right

	return nextVal
}

// HasNext returns whether we have a next smallest number
func (this *BSTIterator) HasNext() bool {
	return !this.st.isEmpty() || this.curNode != nil
}

type stack struct {
	data []*TreeNode
}

func (st *stack) push(v *TreeNode) {
	st.data = append(st.data, v)
}
func (st *stack) pop() *TreeNode {
	v := st.data[len(st.data)-1]
	st.data = st.data[:len(st.data)-1]
	return v
}
func (st *stack) isEmpty() bool {
	return len(st.data) == 0
}
