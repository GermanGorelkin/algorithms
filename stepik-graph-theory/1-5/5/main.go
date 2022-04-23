package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	list := ReadAdjacencyList(N, M)
	used := make([]int, N)

	var U, V int
	fmt.Scanf("%d %d ", &U, &V)

	var path []int
	dfs(U-1, list, used, &path, V)

	if len(path) > 0 {
		for _, v := range path {
			fmt.Printf("%d ", v)
		}
	} else {
		fmt.Printf("no solution")
	}
	fmt.Printf("\n")
}

func dfs(v int, list [][]int, used []int, path *[]int, x int) bool {
	if used[v] != 0 {
		return false
	}

	used[v] = 1
	*path = append(*path, v+1)
	if v+1 == x {
		return true
	}

	for _, v2 := range list[v] {
		if dfs(v2, list, used, path, x) {
			return true
		}
	}

	*path = (*path)[:len(*path)-1]

	return false
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

	// sort
	for i := range list {
		sort.Ints(list[i])
	}

	return list
}
