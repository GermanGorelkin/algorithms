/*
Задача на программирование: покрыть отрезки точками

По данным nn отрезкам необходимо найти множество точек минимального размера,
для которого каждый из отрезков содержит хотя бы одну из точек.
В первой строке дано число 1≤n≤100 отрезков. Каждая из последующих nn строк содержит по два числа 0<=l<=r<=10^9,
задающих начало и конец отрезка. Выведите оптимальное число mm точек и сами mm точек.
Если таких множеств точек несколько, выведите любое из них.

Sample Input 1:
3
1 3
2 5
3 6
Sample Output 1:
1
3

Sample Input 2:
4
4 7
1 3
2 5
5 6
Sample Output 2:
2
3 6

*/

package main

import (
	"fmt"
	"sort"
)

type line struct {
	p1 int
	p2 int
}

func main() {
	var count int
	_, _ = fmt.Scanf("%d", &count)

	lines := make([]line, count)
	for i := 0; i < count; i++ {
		var line line
		_, _ = fmt.Scanf("%d %d", &line.p1, &line.p2)
		lines[i] = line
	}

	points := cover(lines)
	l := len(points)
	fmt.Printf("%d\n", l)
	for _, p := range points {
		fmt.Printf("%d\n", p)
	}
}

func cover(lines []line) []int {
	var points []int
	sort.SliceStable(lines, func(i, j int) bool { return lines[i].p2 < lines[j].p2 })

	p := -1
	for i := 0; i < len(lines); i++ {
		if p < lines[i].p1 {
			p = lines[i].p2
			points = append(points, p)
		}
	}

	return points
}
