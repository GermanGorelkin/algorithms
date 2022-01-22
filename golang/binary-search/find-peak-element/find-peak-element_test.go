package find_peak_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findPeakElement(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"[0 1 2 3 4]":     {input: []int{1, 2, 3, 4, 5}, want: 4},
		"[4 3 2 1 0]":     {input: []int{4, 3, 2, 1, 0}, want: 0},
		"[0 1 2 1 0]":     {input: []int{0, 1, 2, 1, 0}, want: 2},
		"[1,2,3,1]":       {input: []int{1, 2, 3, 1}, want: 2},
		"[1,2,1,3,5,6,4]": {input: []int{1, 2, 1, 3, 5, 6, 4}, want: 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findPeakElement(tc.input))
		})
	}
}
