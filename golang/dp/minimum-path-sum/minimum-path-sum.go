/*
Given a m x n grid filled with non-negative numbers, find a path from top left to bottom right which minimizes the sum of all numbers along its path.

Note: You can only move either down or right at any point in time.

Example:

Input:
[
  [1,3,1],
  [1,5,1],
  [4,2,1]
]
Output: 7
Explanation: Because the path 1→3→1→1→1 minimizes the sum.
*/

package minimum_path_sum

const maxInt = int(^uint(0) >> 1)

func minPathSum(grid [][]int) int {
	h, w := len(grid), len(grid[0])

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			cost1, cost2 := maxInt, maxInt
			if i >= 1 {
				cost1 = grid[i-1][j]
			}
			if j >= 1 {
				cost2 = grid[i][j-1]
			}
			if cost1 != maxInt || cost2 != maxInt {
				grid[i][j] = min(cost1, cost2) + grid[i][j]
			}
		}
	}

	return grid[h-1][w-1]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
