/*
Given a 2d grid map of '1's (land) and '0's (water), count the number of islands. An island is surrounded by water and is formed by connecting adjacent lands horizontally or vertically. You may assume all four edges of the grid are all surrounded by water.

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
	h := len(grid)
	w := len(grid[0])

	q := queue{}
	visited[start] = struct{}{}
	q.enQueue(start)

	for !q.isEmpty() {
		p := q.deQueue()

		var next point

		next = point{i: p.i - 1, j: p.j}
		if next.i >= 0 && grid[next.i][next.j] == 1 && !hasVisited(next, visited) {
			visited[next] = struct{}{}
			q.enQueue(next)
		}
		next = point{i: p.i + 1, j: p.j}
		if next.i < h && grid[next.i][next.j] == 1 && !hasVisited(next, visited) {
			visited[next] = struct{}{}
			q.enQueue(next)
		}

		next = point{i: p.i, j: p.j - 1}
		if next.j >= 0 && grid[next.i][next.j] == 1 && !hasVisited(next, visited) {
			visited[next] = struct{}{}
			q.enQueue(next)
		}
		next = point{i: p.i, j: p.j + 1}
		if next.j < w && grid[next.i][next.j] == 1 && !hasVisited(next, visited) {
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

//---------V2

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
