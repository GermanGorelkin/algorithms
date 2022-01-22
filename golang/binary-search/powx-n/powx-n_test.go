package powx_n

import (
	"testing"
)

func Test_myPow(t *testing.T) {
	type input struct {
		x float64
		n int
	}
	tests := map[string]struct {
		input input
		want  float64
	}{
		"1": {input: input{
			x: 2,
			n: 10,
		}, want: 1024},
		"2": {input: input{
			x: 2.1,
			n: 3,
		}, want: 9.261},
		"3": {input: input{
			x: 2,
			n: -2,
		}, want: 0.25},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := myPow(tc.input.x, tc.input.n)
			if !floatEquals(tc.want, got) {
				t.Errorf("want:%f, got:%f", tc.want, got)
			}
		})
	}
}

var EPSILON float64 = 0.0001

func floatEquals(a, b float64) bool {
	if (a-b) < EPSILON && (b-a) < EPSILON {
		return true
	}
	return false
}
