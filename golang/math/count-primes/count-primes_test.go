package count_primes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_countPrimes(t *testing.T) {
	tests := map[string]struct {
		args int
		want int
	}{
		"10": {args: 10, want: 4},
		"0":  {args: 0, want: 0},
		"1":  {args: 1, want: 0},
		"2":  {args: 2, want: 0},
		"4":  {args: 4, want: 2},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, countPrimes(tc.args))
		})
	}
}
