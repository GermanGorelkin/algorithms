package fibonacci_number

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"1":  {input: 1, want: 1},
		"2":  {input: 2, want: 1},
		"3":  {input: 3, want: 2},
		"4":  {input: 4, want: 3},
		"5":  {input: 5, want: 5},
		"6":  {input: 6, want: 8},
		"10": {input: 10, want: 55},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, fib(tc.input))
		})
	}
}

func Test_frec(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"1":  {input: 1, want: 1},
		"2":  {input: 2, want: 1},
		"3":  {input: 3, want: 2},
		"4":  {input: 4, want: 3},
		"5":  {input: 5, want: 5},
		"6":  {input: 6, want: 8},
		"10": {input: 10, want: 55},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, frec(tc.input))
		})
	}
}
