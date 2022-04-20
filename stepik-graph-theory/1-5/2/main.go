package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	if N == 0 {
		fmt.Printf("%d\n", 0)
		return
	}

	list := ReadAdjacencyList(N, M)
	used := make([]int, N)

	var count int
	for {
		var exists bool
		for i, v := range used {
			if v == 0 {
				exists = true
				count++
				dfs(count, i, list, used)
			}
		}
		if !exists {
			break
		}
	}

	fmt.Printf("%d\n", count)

	comp := make([][]int, count)
	for i, v := range used {
		comp[v-1] = append(comp[v-1], i+1)
	}

	for _, vv := range comp {
		fmt.Printf("%d\n", len(vv))

		for _, v := range vv {
			fmt.Printf("%d ", v)
		}
		fmt.Printf("\n")
	}
}

func dfs(con int, v int, list [][]int, used []int) {
	if used[v] != 0 {
		return
	}

	used[v] = con

	for _, v2 := range list[v] {
		dfs(con, v2, list, used)
	}
}

func ReadAdjacencyList(n, m int) [][]int {
	list := make([][]int, n)

	for i := 0; i < m; i++ {
		var from, to int
		fmt.Scanf("%d %d ", &from, &to)
		from--
		to--

		list[from] = append(list[from], to)
		list[to] = append(list[to], from)
	}

	return list
}
