package find_minimum_in_rotated_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMin(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{3, 4, 5, 1, 2}, want: 1},
		"2": {input: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
		"3": {input: []int{0, 1, 2, 3, 4, 5, 6}, want: 0},
		"4": {input: []int{3, 1}, want: 1},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findMin(tc.input))
		})
	}
}
