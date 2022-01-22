package rotate_array

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_rotate(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		want  []int
		input args
	}{
		"1": {input: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		"2": {input: args{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rotate(tc.input.nums, tc.input.k)
			assert.Equal(t, tc.want, tc.input.nums)
		})
	}
}

func Test_rotate2(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		want  []int
		input args
	}{
		"1": {input: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		"2": {input: args{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rotate2(tc.input.nums, tc.input.k)
			assert.Equal(t, tc.want, tc.input.nums)
		})
	}
}

func Test_rotate3(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		want  []int
		input args
	}{
		"1": {input: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		"2": {input: args{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rotate3(tc.input.nums, tc.input.k)
			assert.Equal(t, tc.want, tc.input.nums)
		})
	}
}

func Test_rotate4(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		want  []int
		input args
	}{
		"1": {input: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		"2": {input: args{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rotate4(tc.input.nums, tc.input.k)
			assert.Equal(t, tc.want, tc.input.nums)
		})
	}
}

func Test_rotate5(t *testing.T) {
	type args struct {
		nums []int
		k    int
	}
	tests := map[string]struct {
		want  []int
		input args
	}{
		"1": {input: args{nums: []int{1, 2, 3, 4, 5, 6, 7}, k: 3}, want: []int{5, 6, 7, 1, 2, 3, 4}},
		"2": {input: args{nums: []int{-1, -100, 3, 99}, k: 2}, want: []int{3, 99, -1, -100}},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			rotate5(tc.input.nums, tc.input.k)
			assert.Equal(t, tc.want, tc.input.nums)
		})
	}
}
