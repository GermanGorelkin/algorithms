package main

import (
	"fmt"
)

type node struct {
	frequency   int
	ch          byte
	right, left *node
}

func createCodeMap(root *node, code string, m map[byte]string) {
	if root.left == nil && root.right == nil {
		if code == "" {
			code = "0"
		}
		m[root.ch] = code
		return
	}
	if root.left != nil {
		createCodeMap(root.left, code+"0", m)
	}
	if root.right != nil {
		createCodeMap(root.right, code+"1", m)
	}
}

func frequency(s string) map[byte]int {
	m := make(map[byte]int)
	for i := 0; i < len(s); i++ {
		m[s[i]]++
	}
	return m
}

type pqueue struct {
	data []*node
}

func (pq *pqueue) insert(val *node) {
	pq.data = append(pq.data, val)
}
func (pq *pqueue) extractMin() *node {
	indx := findMin(pq.data)
	if indx == -1 {
		return nil
	}
	min := pq.data[indx]

	// del indx
	copy(pq.data[indx:], pq.data[indx+1:])
	pq.data = pq.data[:len(pq.data)-1]

	return min
}

func newPQueue(m map[byte]int) *pqueue {
	pq := &pqueue{}
	for ch, f := range m {
		pq.insert(&node{ch: ch, frequency: f})
	}
	return pq
}

func findMin(data []*node) int {
	if len(data) == 0 {
		return -1
	}
	min, minIndx := data[0].frequency, 0
	for i := 1; i < len(data); i++ {
		if data[i].frequency < min {
			min, minIndx = data[i].frequency, i
		}
	}
	return minIndx
}

func createHTree(pq *pqueue) *node {
	for len(pq.data) >= 2 {
		f1 := pq.extractMin()
		f2 := pq.extractMin()

		root := &node{
			frequency: f1.frequency + f2.frequency,
			right:     f2,
			left:      f1,
		}
		pq.insert(root)
	}
	return pq.data[0]
}

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)

	fr := frequency(s)
	pq := newPQueue(fr)
	tree := createHTree(pq)
	codeMap := make(map[byte]string)
	createCodeMap(tree, "", codeMap)

	var code string
	for i := 0; i < len(s); i++ {
		code += codeMap[s[i]]
	}

	fmt.Printf("%d %d\n", len(codeMap), len(code))
	for k, v := range codeMap {
		fmt.Printf("%s: %s\n", string(k), v)
	}
	fmt.Printf("%s", code)
}
