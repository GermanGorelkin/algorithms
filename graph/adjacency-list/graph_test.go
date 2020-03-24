package adjacency_list

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_Graph(t *testing.T) {
	/*

		0----3
		|\
		| \
		|  \
		1---2

	*/
	gr := NewGraph()

	gr.AddVertex(0)
	gr.AddVertex(1)
	gr.AddVertex(2)

	gr.AddEdge(0, 1)
	gr.AddEdge(0, 2)
	gr.AddEdge(0, 3)

	gr.AddEdge(1, 0)
	gr.AddEdge(1, 2)

	gr.AddEdge(2, 0)
	gr.AddEdge(2, 1)

	gr.AddEdge(3, 0)

	want := map[int][]int{
		0: {1, 2, 3},
		1: {0, 2},
		2: {0, 1},
		3: {0},
	}

	assert.Equal(t, want, gr.nodes)
}
