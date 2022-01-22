package remove_element

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type args struct {
	nums []int
	val  int
}
type output struct {
	len  int
	nums []int
}

var tests = map[string]struct {
	input args
	want  output
}{
	"1": {
		input: args{nums: []int{3, 2, 2, 3}, val: 3},
		want:  output{nums: []int{2, 2}, len: 2},
	},
	"2": {
		input: args{nums: []int{0, 1, 2, 2, 3, 0, 4, 2}, val: 2},
		want:  output{nums: []int{0, 1, 3, 0, 4}, len: 5},
	},
}

func Test_removeElement1(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want.len, removeElement1(tc.input.nums, tc.input.val))
			assert.Equal(t, tc.want.nums, tc.want.nums[:tc.want.len])
		})
	}
}

func Test_removeElement2(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want.len, removeElement2(tc.input.nums, tc.input.val))
			assert.Equal(t, tc.want.nums, tc.want.nums[:tc.want.len])
		})
	}
}

func Test_removeElement3(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want.len, removeElement3(tc.input.nums, tc.input.val))
			assert.Equal(t, tc.want.nums, tc.want.nums[:tc.want.len])
		})
	}
}

func Test_removeElement4(t *testing.T) {
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			assert.Equal(t, tc.want.len, removeElement4(tc.input.nums, tc.input.val))
			assert.Equal(t, tc.want.nums, tc.want.nums[:tc.want.len])
		})
	}
}
