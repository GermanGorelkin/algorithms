package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_cover(t *testing.T) {
	tests := map[string]struct {
		input []line
		want  []int
	}{
		"1": {input: []line{{1, 3}, {2, 5}, {3, 6}}, want: []int{3}},
		"2": {input: []line{{4, 7}, {1, 3}, {2, 5}, {5, 6}}, want: []int{3, 6}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, cover(tc.input))
		})
	}
}
