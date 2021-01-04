package kth_largest_element_in_a_stream

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestKthLargest_Add(t *testing.T) {
	kthLargest := Constructor(3, []int{4, 5, 8, 2})
	assert.Equal(t, 4, kthLargest.Add(3))
	assert.Equal(t, 5, kthLargest.Add(5))
	assert.Equal(t, 5, kthLargest.Add(10))
	assert.Equal(t, 8, kthLargest.Add(9))
	assert.Equal(t, 8, kthLargest.Add(4))
}