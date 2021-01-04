package quick_find

type quickFind struct {
	nodes []int
}

func NewQuickFind(n int) *quickFind {
	qf := quickFind{nodes: make([]int, n)}
	for i := range qf.nodes {
		qf.nodes[i] = i
	}
	return &qf
}

func (qf *quickFind) Union(p,q int) {
	pNode := qf.nodes[p]
	qNode := qf.nodes[q]
	for i := range qf.nodes {
		if qf.nodes[i] == qNode {
			qf.nodes[i] = pNode
		}
	}
}

func (qf *quickFind) Connected(p,q int) bool {
	return qf.nodes[p] == qf.nodes[q]
}
