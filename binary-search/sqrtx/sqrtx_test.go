package sqrtx

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_mySqrt(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"4": {input: 4, want: 2},
		"8": {input: 8, want: 2},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, mySqrt(tc.input))
		})
	}
}
