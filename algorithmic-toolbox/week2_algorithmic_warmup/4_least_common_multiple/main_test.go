package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_lcm(t *testing.T) {
	type input struct {
		a, b int
	}
	tests := map[string]struct {
		input input
		want  int
	}{
		"1": {input: input{a: 140, b: 72}, want: 2520},
		"2": {input: input{a: 6, b: 8}, want: 24},
		"3": {input: input{a: 761457, b: 614573}, want: 467970912861},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, lcm(tc.input.a, tc.input.b))
		})
	}
}
