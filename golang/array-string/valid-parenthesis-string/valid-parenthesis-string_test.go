package valid_parenthesis_string

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_checkValidString(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"1":  {input: "()", want: true},
		"2":  {input: "(*)", want: true},
		"3":  {input: "(*))", want: true},
		"4":  {input: "", want: true},
		"5":  {input: "*", want: true},
		"6":  {input: "(())()", want: true},
		"7":  {input: "(", want: false},
		"8":  {input: ")", want: false},
		"9":  {input: "(()", want: false},
		"10": {input: "(()", want: false},
		"11": {input: "(()))", want: false},
		"12": {input: "(*()", want: true},
		"13": {input: "(((******))", want: true},
		"14": {input: "(())((())()()(*)(*()(())())())()()((()())((()))(*", want: false},
	}
	// **((*
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, checkValidString(tc.input))
		})
	}
}
