package union_find

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
// MakeSet creates a singleton set
func (d *DisjointSet) MakeSet(x int) {
	d.data[x] = x
	d.rank[x] = 0
}
// Find finds ID x
func (d *DisjointSet) Find(x int) int {
	for x != d.data[x] {
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
