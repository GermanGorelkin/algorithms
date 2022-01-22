package plus_one

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_plusOne(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"[1,2,3]":   {input: []int{1, 2, 3}, want: []int{1, 2, 4}},
		"[4,3,2,1]": {input: []int{4, 3, 2, 1}, want: []int{4, 3, 2, 2}},
		"[4,3,2,9]": {input: []int{4, 3, 2, 9}, want: []int{4, 3, 3, 0}},
		"[9]":       {input: []int{9}, want: []int{1, 0}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, plusOne(tc.input))
		})
	}
}
