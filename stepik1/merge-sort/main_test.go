package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_sort(t *testing.T) {
	data := []int{10, 9, 5, 4, 1}

	got := sort(data)

	want := []int{1, 4, 5, 9, 10}

	assert.Equal(t, want, got)
}

func Test_sortIter(t *testing.T) {
	data := []int{10, 9, 5, 4, 1}

	got := sortIter(data)

	want := []int{1, 4, 5, 9, 10}

	assert.Equal(t, want, got)
}
