package main

import (
	"fmt"
)

var (
	P, K   uint64
	used   []int
	cycles int
)

func main() {
	fmt.Scanf("%d %d ", &P, &K)
	used = make([]int, P+1)

	for i := range used {
		if i > 0 && used[i] == 0 {
			dfs(i)
		}
	}

	fmt.Printf("%d\n", cycles)
}

func dfs(v int) {
	used[v] = 1

	to := uint64(v) * K % P

	if to > 0 && to <= P {
		if used[to] == 1 {
			cycles++
		} else {
			dfs(int(to))
		}
	}

	used[v] = 2
}
