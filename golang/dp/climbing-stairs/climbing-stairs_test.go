package climbing_stairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_climbStairs(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"1": {input: 1, want: 1},
		"2": {input: 2, want: 2},
		"3": {input: 3, want: 3},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, climbStairs(tc.input))
		})
	}
}
