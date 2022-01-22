package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuickUnion(t *testing.T) {
	qu := NewQuickUnion(10)
	assert.Equal(t, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}, qu.nodes)
}

func Test_quickUnion_Connected(t *testing.T) {
	t.Skip()
}

func Test_quickUnion_Union(t *testing.T) {
	type args struct {
		p int
		q int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// 0, 1, 2, 3, 4, 5, 6, 7, 8, 9
		{name: "0-2", args: args{p: 0, q: 2}, want: []int{0, 1, 0, 3, 4, 5, 6, 7, 8, 9}},
		{name: "2-3", args: args{p: 2, q: 3}, want: []int{0, 1, 0, 0, 4, 5, 6, 7, 8, 9}},
		{name: "7-9", args: args{p: 7, q: 9}, want: []int{0, 1, 0, 0, 4, 5, 6, 7, 8, 7}},
		{name: "6-8", args: args{p: 6, q: 8}, want: []int{0, 1, 0, 0, 4, 5, 6, 7, 6, 7}},
		{name: "4-6", args: args{p: 4, q: 6}, want: []int{0, 1, 0, 0, 6, 5, 6, 7, 6, 7}},
		{name: "8-9", args: args{p: 8, q: 9}, want: []int{0, 1, 0, 0, 6, 5, 6, 6, 6, 7}},
	}

	qu := NewQuickUnion(10)

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			qu.Union(tc.args.p, tc.args.q)
			assert.Equal(t, tc.want, qu.nodes)
		})
	}
}

func Test_quickUnion_root(t *testing.T) {
	qu := quickUnion{nodes: []int{0, 1, 0, 0, 4, 5, 4, 4, 6, 7}}

	tests := []struct {
		name string
		args int
		want int
	}{
		{name: "0", args: 0, want: 0},
		{name: "1", args: 1, want: 1},
		{name: "2", args: 2, want: 0},
		{name: "3", args: 3, want: 0},
		{name: "4", args: 4, want: 4},
		{name: "5", args: 5, want: 5},
		{name: "6", args: 6, want: 4},
		{name: "7", args: 7, want: 4},
		{name: "8", args: 8, want: 4},
		{name: "9", args: 9, want: 4},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.want, qu.root(tc.args))
		})
	}
}

func Test_quickUnion_root2(t *testing.T) {
	/*
							0
						  /
						1
					  /
					2
				  /    \
				4		3
		      /   \
			6		5
		  /
		7
	*/
	qu := quickUnion{nodes: []int{0, 0, 1, 2, 2, 4, 4, 6}}

	assert.Equal(t, 0, qu.root2(6))
	/*
		                    0
		                 /     \
						2	    1
			          / |  \
					6	4	3
			      /     |
				7		5
	*/
	assert.Equal(t, []int{0, 0, 0, 2, 2, 4, 2, 6}, qu.nodes)
}

func Test_quickUnion_root3(t *testing.T) {
	/*
							0
						  /
						1
					  /
					2
				  /    \
				4		3
		      /   \
			6		5
		  /
		7
	*/
	qu := quickUnion{nodes: []int{0, 0, 1, 2, 2, 4, 4, 6}}

	assert.Equal(t, 0, qu.root3(6))
	/*
			                       0
			                 /  |   \    \
							2	1    4    6
		                   /         |    |
					      3          5    7

	*/
	//                    0  1  2  3  4  5  6  7
	assert.Equal(t, []int{0, 0, 0, 2, 0, 4, 0, 6}, qu.nodes)
}