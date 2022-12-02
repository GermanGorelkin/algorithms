package quick

import (
	v1 "github.com/germangorelkin/algorithms/golang/sort/quick/v1"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_v1_qsort(t *testing.T) {
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
		"7": {input: []int{5, 3, 1, 8, 5, 3, 1, 8, 5, 3, 1, 8, 5, 3, 1, 8},
			want: []int{1, 1, 1, 1, 3, 3, 3, 3, 5, 5, 5, 5, 8, 8, 8, 8}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			v1.QuickSort(tc.input, 0, len(tc.input)-1)
			assert.Equal(t, tc.want, tc.input)
		})
	}
}
