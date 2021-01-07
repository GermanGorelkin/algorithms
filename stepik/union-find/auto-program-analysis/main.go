package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, E, D int
	_, _ = fmt.Scanf("%d %d %d", &N, &E, &D)

	if E == 0 || D == 0 {
		fmt.Println(1)
		return
	}

	dset := NewDisjointSet(N)

	reader := bufio.NewReaderSize(os.Stdin, 10_000)
	var line []byte
	var ss []string

	for i := 0; i < E; i++ {
		line, _ = reader.ReadSlice('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		ss = strings.Split(string(line), " ")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])

		dset.Union(x-1, y-1)
	}

	for i := 0; i < D; i++ {
		line, _ = reader.ReadSlice('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		ss = strings.Split(string(line), " ")
		x, _ := strconv.Atoi(ss[0])
		y, _ := strconv.Atoi(ss[1])

		if dset.Connected(x-1, y-1) {
			fmt.Println(0)
			return
		}
	}

	fmt.Println(1)
}



// DisjointSet represents a disjoint-set data structure
// aka union–find data structure, merge–find set, disjoint set forest, disjoint set union.
type DisjointSet struct{
	data []int
	rank []int
}
// NewDisjointSet creates new DisjointSet with n elements
func NewDisjointSet(n int) *DisjointSet {
	ds := DisjointSet{
		data: make([]int, n),
		rank: make([]int, n),
	}
	for i := range ds.data {
		ds.data[i] = i
	}
	return &ds
}
func (d *DisjointSet) Connected(x,y int) bool {
	return d.Find(x) == d.Find(y)
}
// Find finds ID x
func (d *DisjointSet) Find(x int) int {
	if x != d.data[x] {
		d.data[x] = d.Find(d.data[x])
	}
	return d.data[x]
}
// Union union x and y
func (d *DisjointSet) Union(x, y int) {
	xID, yID := d.Find(x), d.Find(y)

	if xID == yID {
		return
	}

	if d.rank[xID] > d.rank[yID] {
		d.data[yID] = xID
	} else {
		d.data[xID] = yID

		if d.rank[xID] == d.rank[yID] {
			d.rank[yID]++
		}
	}
}