package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var N, M int
	_, _ = fmt.Scanf("%d %d", &N, &M)

	writer := bufio.NewWriterSize(os.Stdout, 1000_000)
	reader := bufio.NewReaderSize(os.Stdin, 1000_000)
	var line []byte
	var ss []string

	dset := NewDisjointSet(N)
	var maxRows int

	//
	line, _ = reader.ReadSlice('\n')
	if line[len(line)-1] == '\n' {
		line = line[:len(line)-1]
	}
	ss = strings.Split(string(line), " ")
	for i := 0; i < N; i++ {
		rows, _ := strconv.Atoi(string(ss[i]))
		dset.SetRows(i, rows)

		maxRows = max(maxRows, rows)
	}

	//
	for i := 0; i < M; i++ {
		line, _ = reader.ReadSlice('\n')
		if line[len(line)-1] == '\n' {
			line = line[:len(line)-1]
		}
		ss = strings.Split(string(line), " ")
		dest, _ := strconv.Atoi(ss[0])
		src, _ := strconv.Atoi(ss[1])

		id := dset.Union(dest-1, src-1)
		maxRows = max(dset.GetRows(id), maxRows)

		writer.WriteString(strconv.Itoa(maxRows))
		writer.WriteString("\n")
	}
	writer.Flush()
}

func max(a,b int) int {
	if a > b {
		return a
	}
	return b
}

type table struct {
	rows    int
	symlink int
}

// DisjointSet represents a disjoint-set data structure
// aka union–find data structure, merge–find set, disjoint set forest, disjoint set union.
type DisjointSet struct {
	data []table
	rank []int
}
// NewDisjointSet creates new DisjointSet with n elements
func NewDisjointSet(n int) *DisjointSet {
	ds := DisjointSet{
		data: make([]table, n),
		rank: make([]int, n),
	}
	for i := range ds.data {
		ds.data[i].symlink = i
	}
	return &ds
}
// SetRows sets rows for tableID
func (d *DisjointSet) SetRows(tableID int, rows int) {
	d.data[tableID].rows = rows
}
// SetRows returns rows for tableID
func (d *DisjointSet) GetRows(tableID int) int {
	return d.data[tableID].rows
}
// Find finds ID x
func (d *DisjointSet) Find(x int) int {
	if x != d.data[x].symlink {
		d.data[x].symlink = d.Find(d.data[x].symlink)
	}
	return d.data[x].symlink
}
// Union union x and y and returns new ID
func (d *DisjointSet) Union(dest, src int) (newID int) {
	destID, srcID := d.Find(dest), d.Find(src)

	if destID == srcID {
		newID = destID
		return newID
	}

	if d.rank[destID] > d.rank[srcID] {
		d.data[srcID].symlink = destID
		d.data[destID].rows += d.data[srcID].rows
		d.data[srcID].rows = 0
		newID = destID
	} else {
		d.data[destID].symlink = srcID
		d.data[srcID].rows += d.data[destID].rows
		d.data[destID].rows = 0
		newID = srcID

		if d.rank[destID] == d.rank[srcID] {
			d.rank[srcID]++
		}
	}
	return newID
}