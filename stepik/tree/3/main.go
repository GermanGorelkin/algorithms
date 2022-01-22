/*
https://stepik.org/lesson/45970/step/3?unit=24123
Проверка более общего свойства дерева поиска
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	if n == 0 {
		fmt.Println("CORRECT")
		return
	}

	root := treeBuild(n, os.Stdin)
	if isValidBST(root, nil, nil) {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("INCORRECT")

	}
}

func isValidBST(root, low, high *node) bool {
	if root == nil {
		return true
	}

	if low != nil && root.key < low.key || high != nil && root.key >= high.key {
		return false
	}

	return isValidBST(root.left, low, root) && isValidBST(root.rigth, root, high)
}

func treeBuild(n int, r io.Reader) *node {
	reader := bufio.NewReaderSize(r, 1024)

	var key, left, right int
	nodes := make([]*node, n)
	root := new(node)
	nodes[0] = root

	for i := 0; i < n; i++ {
		fmt.Fscanln(reader, &key, &left, &right)

		if nodes[i] == nil {
			nodes[i] = new(node)
		}
		nodes[i].key = key

		if left != -1 {
			if nodes[left] == nil {
				nodes[left] = new(node)
			}
			nodes[i].left = nodes[left]
		}

		if right != -1 {
			if nodes[right] == nil {
				nodes[right] = new(node)
			}
			nodes[i].rigth = nodes[right]
		}
	}

	return root
}

type node struct {
	key         int
	left, rigth *node
}
