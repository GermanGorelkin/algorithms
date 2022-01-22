package max_consecutive_ones

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findMaxConsecutiveOnes(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{1, 1, 0, 1, 1, 1}, want: 3},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findMaxConsecutiveOnes(tc.input))
		})
	}
}
