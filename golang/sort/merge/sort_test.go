package merge

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMergeSort(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  []int
	}{
		"1": {input: []int{}, want: []int{}},
		"2": {input: []int{1}, want: []int{1}},
		"3": {input: []int{1, 2}, want: []int{1, 2}},
		"4": {input: []int{2, 1}, want: []int{1, 2}},
		"5": {input: []int{2, 1, 3}, want: []int{1, 2, 3}},
		"6": {input: []int{2, 1, 3, 9, 8, 7, 6, 5, 4}, want: []int{1, 2, 3, 4, 5, 6, 7, 8, 9}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, MergeSort(tc.input))
		})
	}
}
