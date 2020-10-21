package main

import "fmt"

/*
4 -1 4 1 1

       1
     /   \
	3	  4
         / \
	    0   2

 */
func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	tree := make(map[int][]int)

	var parent, root int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &parent)
		if parent == -1 {
			root = i
		} else {
			tree[parent] = append(tree[parent], i)
		}
	}

	height := heightTree(tree, root)

	fmt.Printf("%d", height)
}

func heightTree(tree map[int][]int, parent int) int {
	children := tree[parent]

	height := 1
	for _, v := range children {
		height = max(height, heightTree(tree, v)+1)
	}

	return height
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}