package contiguous_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxLength(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{0, 1}, want: 2},
		"2": {input: []int{0, 1, 0}, want: 2},
		"3": {input: []int{0, 1, 0, 1}, want: 4},
		"4": {input: []int{0, 0, 1, 1}, want: 4},
		"5": {input: []int{0, 0, 1, 0, 0, 0, 1, 1}, want: 6},
		"6": {input: []int{0, 0, 1, 0, 0, 0, 1, 0}, want: 2},
		"7": {input: []int{0, 0, 1, 0, 0, 0, 1, 1}, want: 6},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findMaxLength(tc.input))
		})
	}
}
