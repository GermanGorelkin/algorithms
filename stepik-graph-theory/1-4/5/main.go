package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	vertexes := make([]int, N)
	for i := 0; i < M; i++ {
		var from, to int
		fmt.Scanf("%d %d ", &from, &to)
		vertexes[from-1]++
		vertexes[to-1]++
	}

	var needVertex int
	fmt.Scanf("%d ", &needVertex)

	count := vertexes[needVertex-1]

	fmt.Printf("%d \n", count)
}

/*
* 2 wrong answer
 */
func Resolv1(n, m int) {
	_, index := ReadEdges(n, n, false)

	var needVertex int
	fmt.Scanf("%d ", &needVertex)

	var count int
	for i := needVertex; i < len(index); i++ {
		if index[i] != -1 {
			count = index[i] - index[needVertex-1]
			break
		}
	}

	fmt.Printf("%d \n", count)
}

func ReadEdges(n, m int, isOriented bool) ([]Edge, []int) {
	edges := make([]Edge, 0, m)
	for i := 0; i < m; i++ {
		var e Edge
		fmt.Scanf("%d %d ", &e.From, &e.To)

		edges = append(edges, e)
		if !isOriented {
			e.From, e.To = e.To, e.From
			edges = append(edges, e)
		}
	}

	sort.Slice(edges, func(i, j int) bool {
		if edges[i].From == edges[j].From {
			return edges[i].To < edges[j].To
		}
		return edges[i].From < edges[j].From
	})

	indx := make([]int, n)
	for i := range indx {
		indx[i] = -1
	}

	for i, v := range edges {
		if indx[v.From-1] == -1 {
			indx[v.From-1] = i
		}
	}

	return edges, indx
}

type Edge struct {
	From, To int
}
