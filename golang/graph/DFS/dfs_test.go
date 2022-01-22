package DFS

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPathStack(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		/*
		   1
		    \
		     2
		    /
		   3
		*/

		v1 := &vertex{1, false}
		v2 := &vertex{2, false}
		v3 := &vertex{3, false}

		graph := map[*vertex][]*vertex{
			v1: {v2},
			v2: {v1, v3},
			v3: {v2},
		}

		want := []int{1, 2, 3}
		assert.Equal(t, want, findPathStack(graph, v1, 3))
	})

	t.Run("2", func(t *testing.T) {
		/*
			  1
			/   \
			2    4
			 \  / \
			 3--5  6
			        \
			         7
		*/

		v1 := &vertex{1, false}
		v2 := &vertex{2, false}
		v3 := &vertex{3, false}
		v4 := &vertex{4, false}
		v5 := &vertex{5, false}
		v6 := &vertex{6, false}
		v7 := &vertex{7, false}

		graph := map[*vertex][]*vertex{
			v1: {v2, v4},
			v2: {v1, v3},
			v3: {v2, v5},
			v4: {v1, v5, v6},
			v5: {v3, v4},
			v6: {v4, v7},
			v7: {v6},
		}

		want := []int{1, 4, 6, 7}
		assert.Equal(t, want, findPathStack(graph, v1, 7))
	})

	t.Run("start v2", func(t *testing.T) {
		/*
			  1
			/   \
			2    4
			 \  / \
			 3--5  6
			        \
			         7
		*/

		v1 := &vertex{1, false}
		v2 := &vertex{2, false}
		v3 := &vertex{3, false}
		v4 := &vertex{4, false}
		v5 := &vertex{5, false}
		v6 := &vertex{6, false}
		v7 := &vertex{7, false}

		graph := map[*vertex][]*vertex{
			v1: {v2, v4},
			v2: {v1, v3},
			v3: {v2, v5},
			v4: {v1, v5, v6},
			v5: {v3, v4},
			v6: {v4, v7},
			v7: {v6},
		}

		want := []int{2, 3, 5, 4, 6, 7}
		assert.Equal(t, want, findPathStack(graph, v2, 7))
	})

	t.Run("start v2, target 4", func(t *testing.T) {
		/*
			  1
			/   \
			2    4
			 \  / \
			 3--5  6
			        \
			         7
		*/

		v1 := &vertex{1, false}
		v2 := &vertex{2, false}
		v3 := &vertex{3, false}
		v4 := &vertex{4, false}
		v5 := &vertex{5, false}
		v6 := &vertex{6, false}
		v7 := &vertex{7, false}

		graph := map[*vertex][]*vertex{
			v1: {v2, v4},
			v2: {v1, v3},
			v3: {v2, v5},
			v4: {v1, v5, v6},
			v5: {v3, v4},
			v6: {v4, v7},
			v7: {v6},
		}

		want := []int{2, 3, 5, 4}
		assert.Equal(t, want, findPathStack(graph, v2, 4))
	})

	t.Run("target not found", func(t *testing.T) {
		/*
			  1
			/   \
			2    4
			 \  / \
			 3--5  6
			        \
			         7
		*/

		v1 := &vertex{1, false}
		v2 := &vertex{2, false}
		v3 := &vertex{3, false}
		v4 := &vertex{4, false}
		v5 := &vertex{5, false}
		v6 := &vertex{6, false}
		v7 := &vertex{7, false}

		graph := map[*vertex][]*vertex{
			v1: {v2, v4},
			v2: {v1, v3},
			v3: {v2, v5},
			v4: {v1, v5, v6},
			v5: {v3, v4},
			v6: {v4, v7},
			v7: {v6},
		}

		want := []int(nil)
		assert.Equal(t, want, findPathStack(graph, v2, 10))
	})
}
