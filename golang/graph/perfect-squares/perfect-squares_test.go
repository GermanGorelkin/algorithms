package perfect_squares

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numSquares(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"1": {input: 12, want: 3},
		"2": {input: 13, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, numSquares(tc.input))
		})
	}
}
