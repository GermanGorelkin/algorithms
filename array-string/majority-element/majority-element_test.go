package majority_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{3, 2, 3}, want: 3},
		"2": {input: []int{2, 2, 1, 1, 1, 2, 2}, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, majorityElement(tc.input))
		})
	}
}
