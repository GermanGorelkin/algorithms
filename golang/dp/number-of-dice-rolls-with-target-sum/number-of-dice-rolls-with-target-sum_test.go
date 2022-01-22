package number_of_dice_rolls_with_target_sum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_numRollsToTarget(t *testing.T) {
	type input struct {
		d      int
		f      int
		target int
	}
	tests := map[string]struct {
		input input
		want  int
	}{
		"1": {input: input{
			d:      1,
			f:      6,
			target: 3,
		}, want: 1},
		"2": {input: input{
			d:      2,
			f:      6,
			target: 7,
		}, want: 6},
		"3": {input: input{
			d:      2,
			f:      5,
			target: 10,
		}, want: 1},
		"4": {input: input{
			d:      1,
			f:      2,
			target: 3,
		}, want: 0},
		"5": {input: input{
			d:      30,
			f:      30,
			target: 500,
		}, want: 222616187},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, numRollsToTarget(tc.input.d, tc.input.f, tc.input.target))
		})
	}
}
