package basic

import (
	"errors"
	"strconv"
	"strings"
)

var (
	ErrEmptyStack = errors.New("stack is empty")
)

type Stack struct {
	s []int
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Size() int {
	return len(s.s)
}

func (s *Stack) IsEmpty() bool {
	return len(s.s) == 0
}

func (s *Stack) String() string {
	var buf strings.Builder
	buf.WriteString("[")

	for i, v := range s.s {
		buf.WriteString(strconv.Itoa(v))
		if i != len(s.s)-1 {
			buf.WriteString(", ")
		}
	}

	buf.WriteString("]")
	return buf.String()
}

func (s *Stack) Push(val int) {
	s.s = append(s.s, val)
}

func (s *Stack) Pop() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}

	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return val, nil
}

func (s *Stack) Top() (int, error) {
	if s.IsEmpty() {
		return 0, ErrEmptyStack
	}
	return s.s[len(s.s)-1], nil
}
