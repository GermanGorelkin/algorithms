package DFS

type vertex struct {
	val     int
	visited bool
}

type stack struct {
	s []*vertex
}

func (s *stack) push(v *vertex) {
	s.s = append(s.s, v)
}
func (s *stack) pop() *vertex {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}

func findPathStack(graph map[*vertex][]*vertex, start *vertex, target int) (s []int) {
	st := &stack{}
	start.visited = true
	st.push(start)

	for !st.isEmpty() {
		v := st.pop()
		s = append(s, v.val)
		if v.val == target {
			return s
		}
		for _, vr := range graph[v] {
			if vr.visited {
				continue
			}
			vr.visited = true
			st.push(vr)
		}
	}

	return nil
}
