package v2

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findJudge(t *testing.T) {
	type args struct {
		N     int
		trust [][]int
	}
	tests := map[string]struct {
		input args
		want  int
	}{
		"1": {input: args{N: 2, trust: [][]int{{1, 2}}}, want: 2},
		"2": {input: args{N: 3, trust: [][]int{{1, 3}, {2, 3}}}, want: 3},
		"3": {input: args{N: 3, trust: [][]int{{1, 3}, {2, 3}, {3, 1}}}, want: -1},
		"4": {input: args{N: 3, trust: [][]int{{1, 2}, {2, 3}}}, want: -1},
		"5": {input: args{N: 4, trust: [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}}, want: 3},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findJudge(tc.input.N, tc.input.trust))
		})
	}
}
