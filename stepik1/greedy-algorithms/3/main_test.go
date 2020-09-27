package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxK(t *testing.T) {
	tests := map[string]struct {
		n    int
		want []int
	}{
		"1": {n: 4, want: []int{1, 3}},
		"2": {n: 6, want: []int{1, 2, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxK(tc.n))
		})
	}
}
