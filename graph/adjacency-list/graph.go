package adjacency_list

type Graph struct {
	nodes map[int][]int
}

func NewGraph() Graph {
	return Graph{nodes: make(map[int][]int)}
}

func (gr Graph) AddVertex(v int) {
	if _, ok := gr.nodes[v]; !ok {
		gr.nodes[v] = []int{}
	}
}

func (gr Graph) AddEdge(v int, e int) {
	gr.nodes[v] = append(gr.nodes[v], e)
}
