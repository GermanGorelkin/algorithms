package valid_palindrome

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPalindrome(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"1": {input: "A man, a plan, a canal: Panama", want: true},
		"2": {input: "race a car", want: false},
		"3": {input: "", want: true},
		"4": {input: "abba", want: true},
		"5": {input: "aba", want: true},
		"6": {input: "0P", want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, isPalindrome(tc.input))
		})
	}
}
