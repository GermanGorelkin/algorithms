package expn

import (
	"github.com/germangorelkin/algorithms/stack/basic"
	"strconv"
	"strings"
)

/*
Infix Expression	Prefix Expression	Postfix Expression
A+B					+AB					AB+
A+(B*C)				+A*BC				ABC*+
(A+B)*C				*+ABC				AB+C*
*/

const operators = "+-*/%^"

func PostfixEvaluate(tokens []string) int {
	stk := basic.NewStack()
	for _, tkn := range tokens {
		value, err := strconv.Atoi(tkn)

		//tkn is operand
		if err == nil {
			stk.Push(value)
			continue
		}

		//tkn is operator
		operand2, _ := stk.Pop()
		operand1, _ := stk.Pop()
		switch tkn {
		case "+":
			stk.Push(operand1 + operand2)
		case "-":
			stk.Push(operand1 - operand2)
		case "*":
			stk.Push(operand1 * operand2)
		case "/":
			stk.Push(operand1 / operand2)
		}
	}

	res, _ := stk.Pop()
	return res
}

func InfixToPrefix(expn string) string {
	expn = reverseString(expn)
	expn = replaceParentheses(expn)
	expn = InfixToPostfix(expn)
	expn = reverseString(expn)
	return expn
}

func replaceParentheses(s string) string {
	r := []rune(s)
	for i := range r {
		if r[i] == '(' {
			r[i] = ')'
		} else if r[i] == ')' {
			r[i] = '('
		}
	}

	return string(r)
}

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func InfixToPostfix(expn string) string {
	var out strings.Builder
	stk := new(stack)

	for _, ch := range expn {

		if ch == '(' {
			stk.push(ch)
			continue
		}

		if ch == ')' {
			// Печатаем все что есть на стеке до '('
			for !stk.isEmpty() && stk.top() != '(' {
				out.WriteRune(stk.pop())
			}
			// Удаляем со стека '('
			if !stk.isEmpty() && stk.top() == '(' {
				stk.pop()
			}
			continue
		}

		// Операнды печатаем в том же порядке
		if isOperand(ch) {
			out.WriteRune(ch)
			continue
		}

		if isOperator(ch) {
			// Если уже есть оператор на стеке и если он выше приоритетом,
			// то печатаем его.
			// Для '(' вернет 0.
			for !stk.isEmpty() && precedence(ch) <= precedence(stk.top()) {
				out.WriteRune(stk.pop())
			}
			// Новый оператор всегда на стек
			stk.push(ch)
		}
	}

	//
	for !stk.isEmpty() {
		out.WriteRune(stk.pop())
	}

	return out.String()
}

func isOperand(ch rune) bool {
	return !isOperator(ch)
}
func isOperator(ch rune) bool {
	return strings.ContainsAny(string(ch), operators)
}
func precedence(ch rune) int {
	switch ch {
	case '+', '-':
		return 1
	case '*', '/', '%':
		return 2
	case '^':
		return 3
	}
	return 0
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
