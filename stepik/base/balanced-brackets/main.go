package main

import "fmt"

func main() {
	var s string
	_, _ = fmt.Scanf("%s", &s)

	st := &stack{}

	for i := 0; i < len(s); i++ {
		ch := s[i]
		if ch == '[' || ch == '{' || ch == '(' {
			st.push(i)
			continue
		} else if ch == ']' || ch == '}' || ch == ')' {
			if st.isEmpty() {
				fmt.Printf("%d", i+1)
				return
			}
			top := s[st.pop()]
			if top != brackets[ch] {
				fmt.Printf("%d", i+1)
				return
			}
		}
	}
	if !st.isEmpty() {
		var indx int
		for !st.isEmpty() {
			indx = st.pop()
		}
		fmt.Printf("%d", indx+1)
		return
	}

	fmt.Println("Success")
}

var brackets = map[byte]byte{
	'}':'{',
	']':'[',
	')':'(',
}

type stack struct {
	data []int
}
func (s *stack) push(v int) {
	s.data = append(s.data, v)
}
func (s *stack) top() int{
	return s.data[len(s.data)-1]
}
func (s *stack) pop() int {
	v := s.data[len(s.data)-1]
	s.data = s.data[:len(s.data)-1]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.data) == 0
}
