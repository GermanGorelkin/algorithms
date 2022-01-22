package BFS

type vertex struct {
	val     int
	visited bool
}

type queue struct {
	data []*vertex
}

func (q *queue) enQueue(v *vertex) {
	q.data = append(q.data, v)
}
func (q *queue) deQueue() *vertex {
	v := q.data[0]
	q.data = q.data[1:]
	return v
}
func (q *queue) isEmpty() bool {
	return len(q.data) == 0
}

func findPath(graph map[*vertex][]*vertex, start *vertex, target int) (s []int) {
	q := &queue{}
	start.visited = true
	q.enQueue(start)

	for !q.isEmpty() {
		v := q.deQueue()
		s = append(s, v.val)

		if v.val == target {
			return s
		}

		for _, vr := range graph[v] {
			if !vr.visited {
				vr.visited = true
				q.enQueue(vr)
			}
		}
	}

	return s
}
