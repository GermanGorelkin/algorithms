package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scanf("%d ", &N)

	list := readTree(N)

	var leaf []int

	for i, node := range list {
		if len(node) == 0 {
			leaf = append(leaf, i+1)
		}
	}

	sort.Ints(leaf)

	for _, v := range leaf {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
}

/*
последовательность целых чисел p2..p3..pn,
состоящая из n−1  числа, где pi равно номеру предка вершины с номером i.
*/
func readTree(n int) [][]int {
	list := make([][]int, n)

	// i - child
	// p - parent
	for i := 1; i < n; i++ {
		var p int
		fmt.Scanf("%d", &p)

		p--
		list[p] = append(list[p], i)
	}

	return list
}

func PrintList(list [][]int) {
	fmt.Println()
	for _, l := range list {
		fmt.Printf("%v\n", l)
	}
	fmt.Println()
}
