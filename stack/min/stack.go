package min

import (
	"errors"
	"github.com/germangorelkin/algorithms/stack/basic"
)

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack struct {
	base basic.Stack
	min  basic.Stack
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) IsEmpty() bool {
	return s.base.IsEmpty()
}

func (s *Stack) Size() int {
	return s.base.Size()
}

func (s *Stack) Pop() (int, error) {
	s.min.Pop()
	return s.base.Pop()
}

func (s *Stack) Push(val int) {
	cmin, err := s.min.Top()
	if err == basic.ErrEmptyStack || val < cmin {
		s.min.Push(val)
	} else {
		s.min.Push(cmin)
	}

	s.base.Push(val)
}

func (s *Stack) Top() (int, error) {
	return s.base.Top()
}

func (s *Stack) GetMin() (int, error) {
	return s.min.Top()
}
