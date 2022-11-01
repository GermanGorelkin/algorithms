package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	stack := NewStack()
	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		s, _, _ := reader.ReadLine()

		op := strings.Split(string(s), " ")

		switch op[0] {
		case "push":
			{
				stack.Push(op[1])
				fmt.Println("ok")
			}
		case "pop":
			{
				v := stack.Pop()
				fmt.Println(v)
			}
		case "back":
			{
				v := stack.Back()
				fmt.Println(v)
			}
		case "size":
			{
				size := stack.Size()
				fmt.Println(size)
			}
		case "clear":
			{
				stack.Clear()
				fmt.Println("ok")
			}
		case "exit":
			{
				fmt.Println("bye")
				break loop
			}
		}
	}
}

type Stack struct {
	data []string
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Size() int {
	return len(s.data)
}

func (s *Stack) Clear() {
	s.data = s.data[:0]
}

func (s *Stack) Push(v string) {
	s.data = append(s.data, v)
}

func (s *Stack) Pop() string {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}

func (s *Stack) Back() string {
	return s.data[len(s.data)-1]
}
