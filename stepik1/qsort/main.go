/*
https://stepik.org/lesson/13249/step/6?unit=3434
Задача на программирование: точки и отрезки
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type PointType int

const (
	LEFT  PointType = iota - 1 // -1
	POINT                      // 0
	RIGHT                      // 1
)

type Point struct {
	x     int
	ptype PointType
	index int
}

func main() {
	var n, m, l, r int
	_, _ = fmt.Scanf("%d %d", &n, &m)
	points := make([]Point, 0, n+m)

	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < n; i++ {
		s, _ := reader.ReadString('\n')
		ss := strings.Split(s[:len(s)-1], " ")
		l, _ = strconv.Atoi(ss[0])
		r, _ = strconv.Atoi(ss[1])

		points = append(points, Point{x: l, ptype: LEFT, index: -1})
		points = append(points, Point{x: r, ptype: RIGHT, index: -1})
	}

	s, _ := reader.ReadString('\n')
	ss := strings.Split(s, " ")
	for i := 0; i < m; i++ {
		l, _ = strconv.Atoi(ss[i])
		points = append(points, Point{x: l, ptype: POINT, index: i})
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i].x != points[j].x {
			return points[i].x < points[j].x
		}
		return points[i].ptype < points[j].ptype
	})

	result := make([]int, m)
	var cnt int
	for _, p := range points {
		if p.ptype == LEFT {
			cnt++
		} else if p.ptype == RIGHT {
			cnt--
		} else if p.ptype == POINT {
			result[p.index] = cnt
		}
	}

	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}
