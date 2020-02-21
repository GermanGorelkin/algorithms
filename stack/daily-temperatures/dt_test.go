package dt

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

var tests = map[string]struct {
	input []int
	want  []int
}{
	"[]":       {input: []int{}, want: []int{}},
	"[50]":     {input: []int{50}, want: []int{0}},
	"[50, 55]": {input: []int{50, 55}, want: []int{1, 0}},
	"[55, 50]": {input: []int{55, 50}, want: []int{0, 0}},
	"[73, 74, 75, 71, 69, 72, 76, 73]": {
		input: []int{73, 74, 75, 71, 69, 72, 76, 73}, want: []int{1, 1, 4, 2, 1, 1, 0, 0},
	},
}

func TestDailyTemperatures(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, dailyTemperatures(tc.input))
		})
	}
}

func TestDailyTemperatures2(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, dailyTemperatures2(tc.input))
		})
	}
}

func TestDailyTemperatures3(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, dailyTemperatures3(tc.input))
		})
	}
}
