/*
Given a string containing only three types of characters: '(', ')' and '*', write a function to check whether this string is valid. We define the validity of a string by these rules:

Any left parenthesis '(' must have a corresponding right parenthesis ')'.
Any right parenthesis ')' must have a corresponding left parenthesis '('.
Left parenthesis '(' must go before the corresponding right parenthesis ')'.
'*' could be treated as a single right parenthesis ')' or a single left parenthesis '(' or an empty string.
An empty string is also valid.
Example 1:
Input: "()"
Output: True
Example 2:
Input: "(*)"
Output: True
Example 3:
Input: "(*))"
Output: True
Note:
The string size will be in the range [1, 100].
 */

package valid_parenthesis_string

func checkValidString(s string) bool {
	st := &stack{}

	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '(':
			{
				if !st.isEmpty() {
					v := st.pop()
					if v == '*' {
						continue
					}
					st.push(v)
				}
				st.push(s[i])
			}
		case ')':
			{
				if st.isEmpty() {
					return false
				}
				_ = st.pop()
			}
		case '*':
			{
				if !st.isEmpty() {
					v := st.pop()
					if v != '*' {
						continue
					}
				}
				st.push(s[i])
			}
		}
	}

	if !st.isEmpty() {
		if st.pop() != '*' {
			return false
		}
	}

	return true
}



type stack struct {
	s []byte
}
func (s *stack) push(v byte) {
	s.s = append(s.s, v)
}
func (s *stack) pop() byte {
	v := s.s[len(s.s)-1]
	s.s = s.s[:len(s.s)-1]
	return v
}
func (s *stack) isEmpty() bool {
	return len(s.s) == 0
}