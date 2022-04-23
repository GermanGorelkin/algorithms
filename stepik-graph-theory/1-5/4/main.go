package main

import "fmt"

func main() {
	var N, K int
	fmt.Scanf("%d %d ", &N, &K)

	if K == 1 {
		fmt.Printf("%d\n", 1)
		return
	}

	count := Solv(N, K)

	fmt.Printf("%d\n", count)
}

func Solv(n, k int) int {
	used := make([]int, n)
	var count int
	for i := 0; i < n; i++ {
		if used[i] == 0 {
			dfs(i, n, k, used)
			count++
		}
	}
	return count
}

func dfs(v int, n int, k int, used []int) {
	if used[v] == 1 {
		return
	}

	used[v] = 1

	for i := v + 1; i <= n; i++ {
		if (i+v+1)%k == 0 && used[i-1] == 0 {
			fmt.Printf("%d %d \n", v+1, i)
			dfs(i-1, n, k, used)
			break
		}
	}
}
