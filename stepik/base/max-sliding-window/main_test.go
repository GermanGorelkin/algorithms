package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_MaxQueue(t *testing.T){
	// 2 7 3 1 5 2 6 2

	q := maxQueue{}
	// 2 7 3 1
	q.enQueue(2)
	q.enQueue(7)
	q.enQueue(3)
	q.enQueue(1)
	assert.Equal(t, 7, q.max())

	// 7 3 1 5
	q.deQueue()
	q.enQueue(5)
	assert.Equal(t, 7, q.max())

	// 3 1 5 2
	q.deQueue()
	q.enQueue(5)

	assert.Equal(t, 5, q.max())

	// 1 5 2 6
	q.deQueue()
	q.enQueue(6)

	assert.Equal(t, 6, q.max())

	// 5 2 6 2
	q.deQueue()
	q.enQueue(2)

	assert.Equal(t, 6, q.max())
}
