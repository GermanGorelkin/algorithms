package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_gcd(t *testing.T) {
	type input struct {
		a, b int
	}
	tests := map[string]struct {
		input input
		want  int
	}{
		"1": {input: input{a: 18, b: 35}, want: 1},
		"2": {input: input{a: 28851538, b: 1183019}, want: 17657},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, gcd(tc.input.a, tc.input.b))
		})
	}
}
