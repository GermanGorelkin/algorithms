package single_element_in_a_sorted_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_singleNonDuplicate(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{1, 1, 2, 3, 3, 4, 4, 8, 8}, want: 2},
		"2": {input: []int{3, 3, 7, 7, 10, 11, 11}, want: 10},
		"3": {input: []int{1, 2, 2}, want: 1},
		"4": {input: []int{1, 1, 2}, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, singleNonDuplicate(tc.input))
		})
	}
}
