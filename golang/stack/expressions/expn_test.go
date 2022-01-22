package expn

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPostfixEvaluate(t *testing.T) {
	tests := map[string]struct {
		tokens []string
		want   int
	}{
		"2 1 + 3 *":                     {tokens: []string{"2", "1", "+", "3", "*"}, want: 9},
		"4 13 5 / +":                    {tokens: []string{"4", "13", "5", "/", "+"}, want: 6},
		"10 6 9 3 + -11 * / * 17 + 5 *": {tokens: []string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}, want: 22},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, PostfixEvaluate(tc.tokens))
		})
	}
}

func TestInfixToPrefix(t *testing.T) {
	tests := map[string]struct {
		expn string
		want string
	}{
		"A+B":              {expn: "A+B", want: "+AB"},
		"A+B+C":            {expn: "A+B+C", want: "+A+BC"},
		"(A+B)*C":          {expn: "(A+B)*C", want: "*+ABC"},
		"A+B*C":            {expn: "A+B*C", want: "+A*BC"},
		"A+(B*C)":          {expn: "A+(B*C)", want: "+A*BC"},
		"x^y/(5*z)+10":     {expn: "x^y/(5*z)+10", want: "+/^xy*5z10"},
		"x^y/(5*(z+z))+10": {expn: "x^y/(5*(z+z))+10", want: "+/^xy*5+zz10"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := InfixToPrefix(tc.expn)
			assert.Equal(t, tc.want, got)
		})
	}
}

func TestReplaceParentheses(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"()":    {input: "()", want: ")("},
		")(":    {input: ")(", want: "()"},
		"1)2(3": {input: "1)2(3", want: "1(2)3"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, replaceParentheses(tc.input))
		})
	}
}

func TestReverseString(t *testing.T) {
	tests := map[string]struct {
		input string
		want  string
	}{
		"empty": {input: "", want: ""},
		"1":     {input: "1", want: "1"},
		"12":    {input: "12", want: "21"},
		"123":   {input: "123", want: "321"},
		"1234":  {input: "1234", want: "4321"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, reverseString(tc.input))
		})
	}
}

func TestInfixToPostfix(t *testing.T) {
	tests := map[string]struct {
		expn string
		want string
	}{
		"A+B":              {expn: "A+B", want: "AB+"},
		"A+B+C":            {expn: "A+B+C", want: "AB+C+"},
		"(A+B)*C":          {expn: "(A+B)*C", want: "AB+C*"},
		"A+B*C":            {expn: "A+B*C", want: "ABC*+"},
		"A+(B*C)":          {expn: "A+(B*C)", want: "ABC*+"},
		"x^y/(5*z)+10":     {expn: "x^y/(5*z)+10", want: "xy^5z*/10+"},
		"x^y/(5*(z+z))+10": {expn: "x^y/(5*(z+z))+10", want: "xy^5zz+*/10+"},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := InfixToPostfix(tc.expn)
			assert.Equal(t, tc.want, got)
		})
	}
}
