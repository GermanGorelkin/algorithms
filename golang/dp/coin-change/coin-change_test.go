package coin_change

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_coinChange(t *testing.T) {
	type input struct {
		coins  []int
		amount int
	}
	tests := map[string]struct {
		input input
		want  int
	}{
		"1": {input: input{coins: []int{1, 2, 5}, amount: 11}, want: 3},
		"2": {input: input{coins: []int{2}, amount: 3}, want: -1},
		"3": {input: input{coins: []int{186, 419, 83, 408}, amount: 6249}, want: 20},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, coinChange(tc.input.coins, tc.input.amount))
		})
	}
}
