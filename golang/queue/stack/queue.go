package queue

type stack struct {
	data []int
}

func (s *stack) push(val int) {
	s.data = append(s.data, val)
}
func (s *stack) pop() int {
	val := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return val
}
func (s *stack) peak() int {
	return s.data[len(s.data)-1]
}
func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}

type queue struct {
	st1 *stack
	st2 *stack
}

func (q *queue) Push(val int) {
	q.st2.push(val)
}

func (q *queue) Pop() int {
	if q.st1.isEmpty() {
		for !q.st2.isEmpty() {
			q.st1.push(q.st2.pop())
		}
	}

	return q.st1.pop()
}

func (q *queue) Peak() int {
	if q.st1.isEmpty() {
		for !q.st2.isEmpty() {
			q.st1.push(q.st2.pop())
		}
	}

	return q.st1.peak()
}

func (q *queue) IsEmpty() bool {
	return q.st1.isEmpty() && q.st2.isEmpty()
}
