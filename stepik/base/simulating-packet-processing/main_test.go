package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_RingBuffer(t *testing.T) {
	rb := NewRingBuffer(3)

	assert.True(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(packet{arrival: 1})
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(packet{arrival: 2})
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(packet{arrival: 3})
	assert.False(t, rb.isEmpty())
	assert.True(t, rb.isFull())

	assert.Equal(t, 1, rb.deQueue().arrival)
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 2, rb.deQueue().arrival)
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 3, rb.deQueue().arrival)
	assert.True(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	rb.enQueue(packet{arrival: 1})
	rb.enQueue(packet{arrival: 2})
	rb.deQueue()
	rb.enQueue(packet{arrival: 3})
	rb.enQueue(packet{arrival: 4})

	assert.Equal(t, []packet{{arrival: 4}, {arrival: 2}, {arrival: 3}}, rb.data)
	assert.False(t, rb.isEmpty())
	assert.True(t, rb.isFull())

	assert.Equal(t, 2, rb.deQueue().arrival)
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 3, rb.deQueue().arrival)
	assert.False(t, rb.isEmpty())
	assert.False(t, rb.isFull())

	assert.Equal(t, 4, rb.deQueue().arrival)
	assert.True(t, rb.isEmpty())
	assert.False(t, rb.isFull())
}
