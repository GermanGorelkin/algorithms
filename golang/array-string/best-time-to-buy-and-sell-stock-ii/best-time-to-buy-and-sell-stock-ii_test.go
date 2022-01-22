package best_time_to_buy_and_sell_stock_ii

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxProfitV1(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{7, 1, 5, 3, 6, 4}, want: 7},
		"2": {input: []int{1, 2, 3, 4, 5}, want: 4},
		"3": {input: []int{7, 6, 4, 3, 1}, want: 0},
		"4": {input: []int{1}, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxProfitV1(tc.input))
		})
	}
}

func Test_maxProfitV2(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"1": {input: []int{7, 1, 5, 3, 6, 4}, want: 7},
		"2": {input: []int{1, 2, 3, 4, 5}, want: 4},
		"3": {input: []int{7, 6, 4, 3, 1}, want: 0},
		"4": {input: []int{1}, want: 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxProfitV2(tc.input))
		})
	}
}
