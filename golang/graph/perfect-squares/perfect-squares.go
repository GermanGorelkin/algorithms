/*
Given a positive integer n, find the least number of perfect square numbers (for example, 1, 4, 9, 16, ...) which sum to n.

Example 1:

Input: n = 12
Output: 3
Explanation: 12 = 4 + 4 + 4.
Example 2:

Input: n = 13
Output: 2
Explanation: 13 = 4 + 9.
 */

package perfect_squares

func numSquares(n int) int {
	q := &queue{}
	visited := make(map[int]struct{})
	var depth int

	q.enQueue(0)
	visited[0] = struct{}{}

	for !q.isEmpty() {
		size := q.size()
		depth++

		for ; size > 0; size-- {
			u := q.deQueue()
			for i := 1; i*i <= n; i++ {
				v := u + i*i
				if v == n {
					return depth
				}
				if v > n {
					break
				}
				if _, ok := visited[v]; !ok {
					q.enQueue(v)
					visited[v] = struct{}{}
				}
			}
		}
	}

	return depth
}

type queue struct {
	data []int
}
func (q *queue) enQueue(v int) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() int {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}
func (q *queue) size() int {
	return len(q.data)
}