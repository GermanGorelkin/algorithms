package detect_capital

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_detectCapitalUse(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"USA":      {input: "USA", want: true},
		"FlaG":     {input: "FlaG", want: false},
		"leetcode": {input: "leetcode", want: true},
		"leeTcode": {input: "leeTcode", want: false},
		"Leetcode": {input: "leetcode", want: true},
		"l":        {input: "l", want: true},
		"L":        {input: "L", want: true},
		"Le":       {input: "Le", want: true},
		"mL":       {input: "mL", want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, detectCapitalUse(tc.input))
		})
	}
}
