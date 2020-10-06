package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_maxCost(t *testing.T) {
	type args struct {
		items  []item
		weight float64
	}
	tests := map[string]struct {
		input args
		want  float64
	}{
		"1": {input: args{items: []item{{60, 20}, {100, 50}, {120, 30}}, weight: 50}, want: 180.000},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, maxCost(tc.input.items, tc.input.weight))
		})
	}
}
