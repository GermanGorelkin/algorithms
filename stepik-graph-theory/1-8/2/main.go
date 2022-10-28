package main

import (
	"fmt"
)

var (
	list   [][]int
	depths []int
)

func main() {
	var N int
	fmt.Scanf("%d ", &N)

	list = make([][]int, N)
	readTree(N)

	depths = make([]int, N)

	dfs(0)

	for _, v := range depths {
		fmt.Printf("%d ", v)
	}

	fmt.Println()
}

func dfs(v int) int {
	depth := 1

	for _, v := range list[v] {
		depth += dfs(v)
	}

	depths[v] = depth

	return depth
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
последовательность целых чисел p2..p3..pn,
состоящая из n−1  числа, где pi равно номеру предка вершины с номером i.
*/
func readTree(n int) [][]int {
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
