package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	if N == 0 {
		fmt.Printf("%d\n", 0)
		return
	}
	if M == 0 {
		fmt.Printf("%d\n", 1)
		fmt.Printf("%d \n", 1)
		return
	}

	list := ReadAdjacencyList(N, M)
	used := make([]int, N)

	dfs(0, list, used)

	var result []int
	for i, v := range used {
		if v == 1 {
			result = append(result, i)
		}
	}

	fmt.Printf("%d", len(result))
	fmt.Printf("\n")
	for _, v := range result {
		fmt.Printf("%d ", v+1)
	}
	fmt.Printf("\n")
}

func dfs(v int, list [][]int, used []int) {
	if used[v] == 1 {
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

		list[from] = append(list[from], to)
		list[to] = append(list[to], from)
	}

	return list
}
