package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	m := make(map[[2]int]int)

	for i := 0; i < M; i++ {
		var from, to int
		fmt.Scanf("%d %d ", &from, &to)

		if to < from {
			to, from = from, to
		}

		m[[2]int{from, to}]++
	}

	var count int
	for _, v := range m {
		if v >= 2 {
			count++
		}
	}

	fmt.Printf("%d \n", count)
}
