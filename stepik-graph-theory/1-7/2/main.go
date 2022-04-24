package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	if M != N-1 {
		fmt.Println("NO")
		return
	}

	list := ReadAdjacencyList(N, M)
	used := make([]int, N)

	isTree := "YES"
	if dfs(0, list, used, -1) {
		isTree = "NO"
	}

	for _, v := range used {
		if v == 0 {
			isTree = "NO"
		}
	}

	fmt.Println(isTree)
}

func dfs(v int, list [][]int, used []int, p int) bool {
	if used[v] == 1 {
		return true
	}

	used[v] = 1

	for _, v2 := range list[v] {
		if v2 == p {
			continue
		}
		if dfs(v2, list, used, v) {
			return true
		}
	}

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

	return list
}
