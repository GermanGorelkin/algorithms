package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_RingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)

	assert.True(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(1)
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(2)
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(3)
	assert.False(t, rb.isEmpty())
	assert.True(t, rb.isFull())

	assert.Equal(t, 1, rb.deQueue())
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 2, rb.deQueue())
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 3, rb.deQueue())
	assert.True(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(1)
	rb.enQueue(2)
	rb.deQueue()
	rb.enQueue(3)
	rb.enQueue(4)

	assert.Equal(t, []int{4, 2, 3}, rb.data)
	assert.False(t, rb.isEmpty())
	assert.True(t, rb.isFull())

	assert.Equal(t, 2, rb.deQueue())
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 3, rb.deQueue())
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 4, rb.deQueue())
	assert.True(t, rb.isEmpty())
	assert.False(t, rb.isFull())
}
