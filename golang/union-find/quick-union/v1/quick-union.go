package v1

type quickUnion struct {
	nodes []int
}

func NewQuickUnion(n int) *quickUnion {
	qu := quickUnion{nodes: make([]int, n)}
	for i := range qu.nodes {
		qu.nodes[i] = i
	}
	return &qu
}

func (qu *quickUnion) Connected(p,q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *quickUnion) Union(p,q int) {
	pRoot := qu.root(p)
	qRoot := qu.root(q)
	qu.nodes[qRoot] = pRoot
}

func (qu *quickUnion) root(i int) int {
	for qu.nodes[i] != i {
		i = qu.nodes[i]
	}
	return i
}
