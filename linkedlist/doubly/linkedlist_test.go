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
