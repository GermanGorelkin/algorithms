/*
Given a 2d grid map of '1's (land) and '0's (water), counting the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

Example 1:

Input:
11110
11010
11000
00000

Output: 1
Example 2:

Input:
11000
11000
00100
00011

Output: 3
*/

package number_of_islands

//-----------------BFS
var offsets = [5]int{0, 1, 0, -1, 0}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var num int
	visited := make(map[point]struct{})

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				p := point{i: i, j: j}
				if hasVisited(p, visited) {
					continue
				}

				num++
				bfs(p, grid, visited)
			}
		}
	}

	return num
}

func bfs(start point, grid [][]byte, visited map[point]struct{}) {
	q := queue{}
	visited[start] = struct{}{}
	q.enQueue(start)

	for !q.isEmpty() {
		p := q.deQueue()

		for i := 0; i < 4; i++ {
			next := point{
				i: p.i + offsets[i],
				j: p.j + offsets[i+1],
			}
			if next.i < 0 || next.i >= len(grid) || next.j < 0 || next.j >= len(grid[0]) {
				continue
			}
			if hasVisited(next, visited) || grid[next.i][next.j] == 0 {
				continue
			}
			visited[next] = struct{}{}
			q.enQueue(next)
		}
	}
}

func hasVisited(p point, visited map[point]struct{}) bool {
	_, ok := visited[p]
	return ok
}

type point struct {
	i, j int
}

type queue struct {
	data []point
}

func (q *queue) enQueue(v point) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() point {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

//---------DFS

func numIslands2(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}
	var num int

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				num++
				dfs(grid, i, j)
			}
		}
	}

	return num
}

func dfs(grid [][]byte, i, j int) {
	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] == 0 {
		return
	}

	grid[i][j] = 0

	dfs(grid, i-1, j)
	dfs(grid, i+1, j)
	dfs(grid, i, j-1)
	dfs(grid, i, j+1)
}
