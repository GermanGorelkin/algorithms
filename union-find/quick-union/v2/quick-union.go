package v1

type quickUnion struct {
	nodes []int
	rank  []int
}

func NewQuickUnion(n int) *quickUnion {
	qu := quickUnion{
		nodes: make([]int, n),
		rank:  make([]int, n),
	}
	for i := range qu.nodes {
		qu.nodes[i] = i
	}
	return &qu
}

func (qu *quickUnion) Connected(p, q int) bool {
	return qu.root(p) == qu.root(q)
}

func (qu *quickUnion) Union(p, q int) {
	pRoot := qu.root(p)
	qRoot := qu.root(q)

	if pRoot == qRoot {
		return
	}

	if qu.rank[qRoot] > qu.rank[pRoot] {
		qu.nodes[pRoot] = qRoot
	} else {
		qu.nodes[qRoot] = pRoot

		if qu.rank[pRoot] == qu.rank[qRoot] {
			qu.rank[pRoot]++
		}
	}
}

//
func (qu *quickUnion) root3(i int) int {
	for qu.nodes[i] != i {
		qu.nodes[i] = qu.root3(qu.nodes[i])
		i = qu.nodes[i]
	}
	return i
}

// path compression
func (qu *quickUnion) root2(i int) int {
	for qu.nodes[i] != i {
		qu.nodes[i] = qu.nodes[qu.nodes[i]]
		i = qu.nodes[i]
	}
	return i
}

func (qu *quickUnion) root(i int) int {
	for qu.nodes[i] != i {
		i = qu.nodes[i]
	}
	return i
}
