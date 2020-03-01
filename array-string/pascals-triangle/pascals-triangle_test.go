package pascals_triangle

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_generate(t *testing.T) {
	tests := map[string]struct {
		input int
		want  [][]int
	}{
		"1": {input: 5, want: [][]int{
			{1},
			{1, 1},
			{1, 2, 1},
			{1, 3, 3, 1},
			{1, 4, 6, 4, 1},
		}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			//generate(tc.input)
			assert.Equal(t, tc.want, generate(tc.input))
		})
	}
}
