package balanced_parentheses

var (
	parenthesesOpen = map[rune]rune{
		'(': ')',
		'[': ']',
		'{': '}',
	}
	parenthesesClose = map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}
)

func isOpen(r rune) bool {
	_, ok := parenthesesOpen[r]
	return ok
}

func getOpen(close rune) rune {
	return parenthesesClose[close]
}

func IsValid(s string) bool {
	var st stack
	for _, v := range s {
		if isOpen(v) {
			st.push(v)
		} else if st.size() > 0 && getOpen(v) == st.top() {
			st.pop()
		} else {
			return false
		}
	}
	return st.size() == 0
}

type stack struct {
	s []rune
}

func (s *stack) size() int {
	return len(s.s)
}

func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}

func (s *stack) push(val rune) {
	s.s = append(s.s, val)
}

func (s *stack) pop() rune {
	val := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return val
}

func (s *stack) top() rune {
	return s.s[len(s.s)-1]
}
