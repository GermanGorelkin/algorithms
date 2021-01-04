package quick_find

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuickFind(t *testing.T) {
	qf := NewQuickFind(10)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, qf.nodes)
}

func Test_quickFind_Connected(t *testing.T) {
	qf := NewQuickFind(10)
	qf.Union(0,2)
	qf.Union(2,3)
	qf.Union(7,9)
	qf.Union(6,8)
	qf.Union(4,6)
	qf.Union(8,9)

	// 0, 1, 0, 0, 4, 5, 4, 4, 4, 4
	assert.True(t, qf.Connected(0,2))
	assert.True(t, qf.Connected(0,3))
	assert.True(t, qf.Connected(2,3))
	assert.True(t, qf.Connected(4,6))
	assert.True(t, qf.Connected(4,9))

	assert.False(t, qf.Connected(0,1))
	assert.False(t, qf.Connected(1,2))
	assert.False(t, qf.Connected(4,5))
}

func Test_quickFind_Union(t *testing.T) {
	qf := NewQuickFind(10)

	type args struct {
		p, q int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "0-2", args: args{0, 2}, want: []int{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}},
		{name: "2-3", args: args{2, 3}, want: []int{0, 1, 0, 0, 4, 5, 6, 7, 8, 9}},
		{name: "7-9", args: args{7, 9}, want: []int{0, 1, 0, 0, 4, 5, 6, 7, 8, 7}},
		{name: "6-8", args: args{6, 8}, want: []int{0, 1, 0, 0, 4, 5, 6, 7, 6, 7}},
		{name: "4-6", args: args{4, 6}, want: []int{0, 1, 0, 0, 4, 5, 4, 7, 4, 7}},
		{name: "8-9", args: args{8, 9}, want: []int{0, 1, 0, 0, 4, 5, 4, 4, 4, 4}},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			qf.Union(tc.args.p, tc.args.q)
			assert.Equal(t, tc.want, qf.nodes)
		})
	}
}