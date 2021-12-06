/*
https://stepik.org/lesson/45970/step/1?unit=24123
Обход дерева
*/

package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	root := treeBuild(n, os.Stdin)

	buf := new(bytes.Buffer)

	inOrder(root, buf)
	fmt.Println(buf.String())
	buf.Reset()
	preOrder(root, buf)
	fmt.Println(buf.String())
	buf.Reset()
	postOrder(root, buf)
	fmt.Println(buf.String())
	buf.Reset()
}

func inOrder(root *node, buf *bytes.Buffer) {
	if root == nil {
		return
	}

	inOrder(root.left, buf)
	buf.WriteString(strconv.Itoa(root.key) + " ")
	inOrder(root.rigth, buf)
}

func preOrder(root *node, buf *bytes.Buffer) {
	if root == nil {
		return
	}

	buf.WriteString(strconv.Itoa(root.key) + " ")
	preOrder(root.left, buf)
	preOrder(root.rigth, buf)
}

func postOrder(root *node, buf *bytes.Buffer) {
	if root == nil {
		return
	}

	postOrder(root.left, buf)
	postOrder(root.rigth, buf)
	buf.WriteString(strconv.Itoa(root.key) + " ")
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
