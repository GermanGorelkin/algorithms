package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	list := ReadAdjacencyList(N, M)
	used := make([]int, N)

	dfs(0, list, used)

	for i, v := range used {
		if v == 1 {
			fmt.Printf("%d ", i+1)
		}
	}
}

func dfs(v int, list [][]int, used []int) {
	if used[v] != 0 {
		return
	}

	used[v] = 1

	for _, v2 := range list[v] {
		dfs(v2, list, used)
	}
}

func ReadAdjacencyList(n, m int) [][]int {
	list := make([][]int, n)

	for i := 0; i < m; i++ {
		var from, to int
		fmt.Scanf("%d %d ", &from, &to)
		from--
		to--

		// reverse
		list[to] = append(list[to], from)
	}

	return list
}
