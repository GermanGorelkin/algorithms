/*
https://stepik.org/lesson/45970/step/2?unit=24123
Проверка свойства дерева поиска
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
	var vl validator
	isBTS := vl.useInOrder(root)
	if isBTS {
		fmt.Println("CORRECT")
	} else {
		fmt.Println("INCORRECT")

	}
}

type validator struct {
	prev *node
}

func (v *validator) useInOrder(root *node) bool {
	if root == nil {
		return true
	}

	if !v.useInOrder(root.left) {
		return false
	}
	if v.prev != nil && v.prev.key > root.key {
		return false
	}
	v.prev = root
	return v.useInOrder(root.rigth)
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
