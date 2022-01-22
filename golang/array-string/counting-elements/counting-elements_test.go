package counting_elements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countElements(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{1, 2, 3}, want: 2},
		"2": {input: []int{1, 1, 3, 3, 5, 5, 7, 7}, want: 0},
		"3": {input: []int{1, 3, 2, 3, 5, 0}, want: 3},
		"4": {input: []int{1, 1, 2, 2}, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, countElements(tc.input))
		})
	}
}
