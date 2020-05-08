/*
Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

Search for a node to remove.
If the node is found, delete the node.
Note: Time complexity should be O(height of tree).

Example:

root = [5,3,6,2,4,null,7]
key = 3

    5
   / \
  3   6
 / \   \
2   4   7

Given key to delete is 3. So we find the node with value 3 and delete it.

One valid answer is [5,4,6,2,null,null,7], shown in the following BST.

    5
   / \
  4   6
 /     \
2       7

Another valid answer is [5,2,6,null,4,null,7].

    5
   / \
  2   6
   \   \
    4   7
*/

package delete_node_in_a_bst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func deleteNode(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return root
	}

	node, parent := search(root, key)

	// not found key
	if node == nil {
		return root
	}

	// has no children
	if node.Right == nil && node.Left == nil {
		if parent == nil {
			return nil
		}

		if parent.Left == node {
			parent.Left = nil
		} else {
			parent.Right = nil
		}
		return root
	}

	// has one child
	if node.Left != nil && node.Right == nil {
		if parent == nil {
			return node.Left
		}

		if parent.Left == node {
			parent.Left = node.Left
		} else {
			parent.Right = node.Left
		}
		return root

	}
	if node.Left == nil && node.Right != nil {
		if parent == nil {
			return node.Right
		}

		if parent.Left == node {
			parent.Left = node.Right
		} else {
			parent.Right = node.Right
		}
		return root
	}

	// has two children
	// Successor node = the node in the right subtree that has the minimum value.
	/*
		1. find min in right subtree
		2. min.left = del.left
		3.1 if parent is null than root = del.right
		3.2 else parent.left\right = del.right
	*/
	min := findMin(node.Right)
	min.Left = node.Left
	if parent == nil {
		return node.Right
	}
	if parent.Left == node {
		parent.Left = node.Right
	} else {
		parent.Right = node.Right
	}

	return root
}

func findMin(root *TreeNode) *TreeNode {
	for root.Left != nil {
		root = root.Left
	}
	return root
}

func search(root *TreeNode, key int) (*TreeNode, *TreeNode) {
	if root == nil || root.Val == key {
		return root, nil
	}

	var node, parent *TreeNode
	if root.Val > key {
		node, parent = search(root.Left, key)
	} else {
		node, parent = search(root.Right, key)
	}

	if node != nil && parent == nil {
		return node, root
	}
	return node, parent
}
