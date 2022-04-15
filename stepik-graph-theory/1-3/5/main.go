package main

import (
	"fmt"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)
	list := make([]edge, M)

	for i := 0; i < M; i++ {
		var e edge
		fmt.Scanf("%d %d", &e.a, &e.b)
		list[i] = e
	}

	var K int
	fmt.Scanf("%d ", &K)
	K--
	e := list[K]

	var result int
	for i, v := range list {
		if i != K && (v.a == e.a || v.a == e.b || v.b == e.a || v.b == e.b) {
			result++
		}
	}

	fmt.Printf("%d\n", result)
}

type edge struct {
	a, b int
}
