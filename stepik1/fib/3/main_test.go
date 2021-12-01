package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_periodPisano(t *testing.T) {
	tests := []struct {
		name string
		args uint
		want []uint
	}{
		{name: "2", args: 2, want: []uint{0, 1, 1}},
		{name: "3", args: 3, want: []uint{0, 1, 1, 2, 0, 2, 2, 1}},
		{name: "4", args: 4, want: []uint{0, 1, 1, 2, 3, 1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, pisanoPeriod(tt.args))
		})
	}
}

func Test_calc(t *testing.T) {
	type args struct {
		n int
		m int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{name: "9 2", args: args{n: 9, m: 2}, want: 0},
		{name: "10 2", args: args{n: 10, m: 2}, want: 1},
		{name: "1025 55", args: args{n: 1025, m: 55}, want: 5},
		{name: "12589 369", args: args{n: 12589, m: 369}, want: 89},
		{name: "1598753 25897", args: args{n: 1598753, m: 25897}, want: 20305},
		{name: "60282445765134413 2263", args: args{n: 60282445765134413, m: 2263}, want: 974},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, calc(tt.args.n, tt.args.m))
		})
	}
}
