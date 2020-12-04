package reverse_integer

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_reverse(t *testing.T) {
	tests := map[string]struct {
		args int
		want int
	}{
		"1":          {args: 1, want: 1},
		"10":         {args: 10, want: 1},
		"100":        {args: 100, want: 1},
		"123":        {args: 123, want: 321},
		"-123":       {args: -123, want: -321},
		"102":        {args: 102, want: 201},
		"1534236469": {args: 1534236469, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, reverse(tc.args))
		})
	}
}
