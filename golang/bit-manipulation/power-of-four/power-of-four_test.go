package power_of_four

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_isPowerOfFour(t *testing.T) {
	tests := map[string]struct {
		input int
		want  bool
	}{
		"16": {input: 16, want: true},
		"8":  {input: 8, want: true},
		"1":  {input: 1, want: true},
		"5":  {input: 5, want: false},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, isPowerOfFour(tc.input))
		})
	}
}
