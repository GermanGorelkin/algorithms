package doubly

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestList_AddHead(t *testing.T) {
	l := NewList()

	l.AddHead(1)
	assert.Equal(t, "[1]", l.String())

	l.AddHead(2)
	assert.Equal(t, "[2, 1]", l.String())

	l.AddHead(3)
	assert.Equal(t, "[3, 2, 1]", l.String())
}
func TestList_Size(t *testing.T) {
	l := NewList()
	assert.Equal(t, 0, l.Size())

	l.AddHead(1)
	assert.Equal(t, 1, l.Size())

	l.AddTail(2)
	assert.Equal(t, 2, l.Size())
}
func TestList_AddTail(t *testing.T) {
	l := NewList()

	l.AddTail(1)
	assert.Equal(t, "[1]", l.String())

	l.AddTail(2)
	assert.Equal(t, "[1, 2]", l.String())

	l.AddTail(3)
	assert.Equal(t, "[1, 2, 3]", l.String())
}
