package find_k_closest_elements

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_findClosestElements(t *testing.T) {
	type input struct {
		arr []int
		k   int
		x   int
	}
	tests := map[string]struct {
		input input
		want  []int
	}{
		"1": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   3,
			},
			want: []int{1, 2, 3, 4},
		},
		"1.1": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   2,
				x:   3,
			},
			want: []int{2, 3},
		},
		"2": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   1,
			},
			want: []int{1, 2, 3, 4},
		},
		"3": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   6,
			},
			want: []int{2, 3, 4, 5},
		},
		"4": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   5,
			},
			want: []int{2, 3, 4, 5},
		},
		"6": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   2,
			},
			want: []int{1, 2, 3, 4},
		},
		"7": {
			input: input{
				arr: []int{1, 2, 3, 4, 5},
				k:   4,
				x:   4,
			},
			want: []int{2, 3, 4, 5},
		},
		"8": {
			input: input{
				arr: []int{1, 2, 4, 5},
				k:   2,
				x:   3,
			},
			want: []int{2, 4},
		},
		"9": {
			input: input{
				arr: []int{1, 1, 1, 10, 10, 10},
				k:   1,
				x:   9,
			},
			want: []int{10},
		},
		"10": {
			input: input{
				arr: []int{1, 1, 1, 10, 10, 10},
				k:   1,
				x:   6,
			},
			want: []int{10},
		},
		"11": {
			input: input{
				arr: []int{1, 1, 1, 10, 10, 10},
				k:   1,
				x:   5,
			},
			want: []int{1},
		},
		"12": {
			input: input{
				arr: []int{0, 0, 0, 1, 3, 5, 6, 7, 8, 8},
				k:   2,
				x:   2,
			},
			want: []int{1, 3},
		},
		"13": {
			input: input{
				arr: []int{1, 2, 5, 5, 6, 6, 7, 7, 8, 9},
				k:   7,
				x:   7,
			},
			want: []int{5, 5, 6, 6, 7, 7, 8},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want, findClosestElements(tc.input.arr, tc.input.k, tc.input.x))
		})
	}
}
