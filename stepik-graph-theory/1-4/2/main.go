package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	Solv2(N, M)
}

/*
*
 */
func Solv2(n, m int) {
	vertexes := make([]int, n)

	for i := 0; i < m; i++ {
		var from, to int
		fmt.Scanf("%d %d ", &from, &to)
		vertexes[from-1]++
		vertexes[to-1]++
	}

	set := make(map[int]struct{})

	for _, v := range vertexes {
		set[v] = struct{}{}
	}

	fmt.Printf("%d \n", len(set))
}

/*
*	решение через упорядоченный список дуг
 */
func Solv1(n, m int) {
	edges, index := ReadEdges(n, m, false)
	index = append(index, len(edges))

	set := make(map[int]struct{})
	prev := -1
	for i := 0; i < len(index); i++ {
		if prev == -1 {
			prev = index[i]
			continue
		}
		if index[i] == -1 {
			set[0] = struct{}{}
			continue
		}

		diff := index[i] - prev
		set[diff] = struct{}{}
		prev = index[i]
	}

	fmt.Printf("%d \n", len(set))
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
