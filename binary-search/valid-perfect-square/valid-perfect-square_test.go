package valid_perfect_square

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPerfectSquare(t *testing.T) {

	tests := map[string]struct {
		input int
		want  bool
	}{
		"16": {input: 16, want: true},
		"14": {input: 14, want: false},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, isPerfectSquare(tc.input))
		})
	}
}
