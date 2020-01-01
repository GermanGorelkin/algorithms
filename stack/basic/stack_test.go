package basic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_String(t *testing.T) {
	s := NewStack()

	t.Run("[]", func(t *testing.T) {
		assert.Equal(t, "[]", s.String())
	})
	t.Run("[1]", func(t *testing.T) {
		s.s = append(s.s, 1)
		assert.Equal(t, "[1]", s.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		s.s = append(s.s, 2)
		assert.Equal(t, "[1, 2]", s.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		s.s = append(s.s, 3)
		assert.Equal(t, "[1, 2, 3]", s.String())
	})
}

func TestStack_Push(t *testing.T) {
	s := NewStack()

	t.Run("[]", func(t *testing.T) {
		assert.Equal(t, "[]", s.String())
	})
	t.Run("[1]", func(t *testing.T) {
		s.Push(1)
		assert.Equal(t, "[1]", s.String())
	})
	t.Run("[1, 2]", func(t *testing.T) {
		s.Push(2)
		assert.Equal(t, "[1, 2]", s.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		s.Push(3)
		assert.Equal(t, "[1, 2, 3]", s.String())
	})
}

func TestStack_Pop(t *testing.T) {
	t.Run("[]", func(t *testing.T) {
		s := NewStack()

		n, err := s.Pop()
		assert.Equal(t, ErrEmptyStack, err)
		assert.Equal(t, 0, n)
		assert.Equal(t, "[]", s.String())
	})
	t.Run("[1, 2, 3]", func(t *testing.T) {
		s := NewStack()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		n, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 3, n)
		assert.Equal(t, "[1, 2]", s.String())

		n, err = s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 2, n)
		assert.Equal(t, "[1]", s.String())

		n, err = s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 1, n)
		assert.Equal(t, "[]", s.String())

		n, err = s.Pop()
		assert.Equal(t, ErrEmptyStack, err)
		assert.Equal(t, 0, n)
		assert.Equal(t, "[]", s.String())
	})
}
