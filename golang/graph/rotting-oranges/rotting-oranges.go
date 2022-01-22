/*
In a given grid, each cell can have one of three values:

the value 0 representing an empty cell;
the value 1 representing a fresh orange;
the value 2 representing a rotten orange.
Every minute, any fresh orange that is adjacent (4-directionally) to a rotten orange becomes rotten.

Return the minimum number of minutes that must elapse until no cell has a fresh orange.  If this is impossible, return -1 instead.
*/

package rotting_oranges

var offsets = [5]int{0, 1, 0, -1, 0}

func orangesRotting(grid [][]int) int {
	h := len(grid)
	w := len(grid[0])
	min := 0
	freshCount := 0
	q := &queue{}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			if grid[i][j] == 2 {
				c := cell{i: i, j: j}
				grid[c.i][c.j] = 0
				q.enQueue(c)
			} else if grid[i][j] == 1 {
				freshCount++
			}
		}
	}

	if freshCount == 0 {
		return 0
	}

	for !q.isEmpty() {
		hasFresh := false

		for l := q.len(); l > 0; l-- {
			c := q.deQueue()

			for i := 0; i < 4; i++ {
				next := cell{
					i: c.i + offsets[i],
					j: c.j + offsets[i+1],
				}
				if next.i < 0 || next.j >= w || next.i >= h || next.j < 0 ||
					grid[next.i][next.j] == 0 || grid[next.i][next.j] == 2 {
					continue
				}
				if grid[next.i][next.j] == 1 {
					freshCount--
					hasFresh = true
				}
				grid[next.i][next.j] = 0
				q.enQueue(next)
			}
		}
		if hasFresh {
			min++
		}
	}

	if freshCount != 0 {
		return -1
	}

	return min
}

type cell struct {
	i, j int
}
type queue struct {
	data []cell
}

func (q *queue) enQueue(v cell) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() cell {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}
func (q *queue) len() int {
	return len(q.data)
}
