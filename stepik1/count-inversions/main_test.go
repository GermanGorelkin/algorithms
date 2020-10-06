package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_countInv(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1":  {input: []int{2, 3, 9, 2, 9}, want: 2},
		"2":  {input: []int{7, 6, 5, 4, 3, 2, 1}, want: 21},
		"3":  {input: []int{1, 2, 3, 5, 4}, want: 1},
		"4":  {input: []int{10, 8, 6, 2, 4, 5}, want: 12},
		"5":  {input: []int{5, 7, 0, 2, 2, 0}, want: 10},
		"6":  {input: []int{9, 10, 9, 5, 7, 7}, want: 10},
		"7":  {input: []int{6, 5, 8, 6, 0, 4}, want: 10},
		"8":  {input: []int{1, 2, 3, 4, 5, 6, 7, 8, 3, 4, 3}, want: 15},
		"9":  {input: []int{1, 3, 4, 5, 6, 2}, want: 4},
		"10": {input: []int{10, 9, 3, 8, 3, 10}, want: 8},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := countInv(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}
