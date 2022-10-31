package main

import (
	"fmt"
)

var (
	list [][]int
	used []int
	root int
	ans  int
)

func main() {
	var N int
	fmt.Scanf("%d ", &N)

	used = make([]int, N)
	list = make([][]int, N)
	readTree(N)
	root = 0

	dfs(root, 0)

	used = make([]int, N)
	dfs(root, 0)

	fmt.Printf("%d \n", ans-1)
}

func dfs(v int, depth int) {
	used[v] = 1
	depth++

	if ans < depth {
		ans = depth
		root = v
	}
	for _, to := range list[v] {
		if used[to] == 0 {
			dfs(to, depth)
		}
	}
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
		list[i] = append(list[i], p)
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
