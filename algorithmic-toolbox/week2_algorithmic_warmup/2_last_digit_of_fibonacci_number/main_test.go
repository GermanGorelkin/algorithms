package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_fib(t *testing.T) {
	tests := map[string]struct {
		input int
		want  int
	}{
		"0":       {input: 0, want: 0},
		"1":       {input: 1, want: 1},
		"2":       {input: 2, want: 1},
		"3":       {input: 3, want: 2},
		"4":       {input: 4, want: 3},
		"5":       {input: 5, want: 5},
		"10":      {input: 10, want: 5},
		"11":      {input: 11, want: 9},
		"20":      {input: 20, want: 5},
		"200":     {input: 200, want: 5},
		"331":     {input: 331, want: 9},
		"327305 ": {input: 327305, want: 5},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, fib(tc.input))
		})
	}
}
