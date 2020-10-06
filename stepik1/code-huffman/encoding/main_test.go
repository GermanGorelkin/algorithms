package main

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func Test_createCodeMap(t *testing.T) {
	/*
			   root
			   /  \
			  18   21
		     /  \
		    10   8
			    / \
			   3   5

			a:00
			b:010
			c:011
			d:1
	*/

	n10 := &node{
		frequency: 10,
		ch:        'a',
	}
	n3 := &node{
		frequency: 3,
		ch:        'b',
	}
	n5 := &node{
		frequency: 5,
		ch:        'c',
	}
	n8 := &node{
		frequency: 8,
		right:     n5,
		left:      n3,
	}
	n18 := &node{
		frequency: 18,
		right:     n8,
		left:      n10,
	}
	n21 := &node{
		frequency: 21,
		ch:        'd',
	}
	root := &node{
		right: n21,
		left:  n18,
	}

	var code string
	m := make(map[byte]string)
	createCodeMap(root, code, m)

	assert.Equal(t, "00", m['a'])
	assert.Equal(t, "010", m['b'])
	assert.Equal(t, "011", m['c'])
	assert.Equal(t, "1", m['d'])
}

func Test_frequency(t *testing.T) {
	s := "abacabad"
	got := frequency(s)
	m := map[byte]int{
		'a': 4,
		'b': 2,
		'c': 1,
		'd': 1,
	}

	for ch, f := range got {
		v, ok := m[ch]
		assert.True(t, ok)
		assert.Equal(t, v, f)
	}
}

func Test_findMin(t *testing.T) {
	data := []*node{
		{frequency: 10},
		{frequency: 2},
		{frequency: 11},
		{frequency: 12},
	}

	got := findMin(data)
	want := 1

	assert.Equal(t, want, got)
}

func Test_pqueue(t *testing.T) {
	pq := pqueue{}

	pq.insert(&node{frequency: 10})
	pq.insert(&node{frequency: 2})
	pq.insert(&node{frequency: 11})
	pq.insert(&node{frequency: 12})

	gotMin := pq.extractMin()
	wantFr := 2
	assert.Equal(t, wantFr, gotMin.frequency)

	wantData := []*node{
		{frequency: 10},
		{frequency: 11},
		{frequency: 12},
	}
	for i, v := range wantData {
		assert.Equal(t, v.frequency, pq.data[i].frequency)
	}
}

func Test_newPQueue(t *testing.T) {
	m := map[byte]int{
		'a': 3,
		'b': 5,
		'c': 10,
		'd': 21,
	}
	pq := newPQueue(m)

	sort.SliceStable(pq.data, func(i, j int) bool {
		return pq.data[i].frequency < pq.data[j].frequency
	})

	want := []node{
		{frequency: 3, ch: 'a'},
		{frequency: 5, ch: 'b'},
		{frequency: 10, ch: 'c'},
		{frequency: 21, ch: 'd'},
	}

	for i, v := range want {
		assert.Equal(t, v.frequency, pq.data[i].frequency)
		assert.Equal(t, v.ch, pq.data[i].ch)
	}
}

func Test_createHTree(t *testing.T) {
	/*
		                39
					   /  \
					  18   21
				     /  \
				    8   10
		           / \
		          3   5
	*/

	m := map[byte]int{
		'a': 3,
		'b': 5,
		'c': 10,
		'd': 21,
	}
	pq := newPQueue(m)

	root := createHTree(pq)

	want := []int{39, 18, 8, 3, 5, 10, 21}
	got := getPreorderTree(root)
	assert.Equal(t, want, got)
}

func getPreorderTree(root *node) (s []int) {
	if root == nil {
		return s
	}
	s = append(s, root.frequency)
	s = append(s, getPreorderTree(root.left)...)
	s = append(s, getPreorderTree(root.right)...)
	return s
}
