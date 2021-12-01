package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fin(t *testing.T) {
	tests := []struct {
		name string
		args int
		want int
	}{
		{name: "1", args: 1, want: 1},
		{name: "2", args: 2, want: 1},
		{name: "3", args: 3, want: 2},
		{name: "4", args: 4, want: 3},
		{name: "5", args: 5, want: 5},
		{name: "10", args: 10, want: 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(*testing.T) {
			assert.Equal(t, tt.want, fib(tt.args))
		})
	}
}
