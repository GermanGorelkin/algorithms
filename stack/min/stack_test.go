package min

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack_Push(t *testing.T) {
	s := NewStack()

	t.Run("base:[5], min:[5]", func(t *testing.T) {
		s.Push(5)

		assert.Equal(t, "[5]", s.base.String())
		assert.Equal(t, "[5]", s.min.String())
	})
	t.Run("base:[5, 4], min:[5, 4]", func(t *testing.T) {
		s.Push(4)

		assert.Equal(t, "[5, 4]", s.base.String())
		assert.Equal(t, "[5, 4]", s.min.String())
	})
	t.Run("base:[5, 4, 6], min:[5, 4, 4]", func(t *testing.T) {
		s.Push(6)

		assert.Equal(t, "[5, 4, 6]", s.base.String())
		assert.Equal(t, "[5, 4, 4]", s.min.String())
	})
}

func TestStack_Pop(t *testing.T) {
	s := NewStack()
	s.Push(4)
	s.Push(3)
	s.Push(6)
	s.Push(2)

	t.Run("base:[4, 3, 6, 2], min:[4, 3, 3, 2]", func(t *testing.T) {
		n, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 2, n)
		assert.Equal(t, "[4, 3, 6]", s.base.String())
		assert.Equal(t, "[4, 3, 3]", s.min.String())
	})
	t.Run("base:[4, 3, 6], min:[4, 3, 3]", func(t *testing.T) {
		n, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 6, n)
		assert.Equal(t, "[4, 3]", s.base.String())
		assert.Equal(t, "[4, 3]", s.min.String())
	})
	t.Run("base:[4, 3], min:[4, 3]", func(t *testing.T) {
		n, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 3, n)
		assert.Equal(t, "[4]", s.base.String())
		assert.Equal(t, "[4]", s.min.String())
	})
	t.Run("base:[4], min:[4]", func(t *testing.T) {
		n, err := s.Pop()
		assert.Nil(t, err)
		assert.Equal(t, 4, n)
		assert.Equal(t, "[]", s.base.String())
		assert.Equal(t, "[]", s.min.String())
	})
	t.Run("base:[], min:[]", func(t *testing.T) {
		_, err := s.Pop()
		assert.Equal(t, ErrEmptyStack, err)
		assert.Equal(t, "[]", s.base.String())
		assert.Equal(t, "[]", s.min.String())
	})
}

func TestStack_GetMin(t *testing.T) {
	s := NewStack()
	s.Push(4)
	s.Push(3)
	s.Push(6)
	s.Push(2)

	t.Run("base:[4, 3, 6, 2], min:[4, 3, 3, 2]", func(t *testing.T) {
		n, err := s.GetMin()
		assert.Nil(t, err)
		assert.Equal(t, 2, n)
	})
	t.Run("base:[4, 3, 6], min:[4, 3, 3]", func(t *testing.T) {
		s.Pop()

		n, err := s.GetMin()
		assert.Nil(t, err)
		assert.Equal(t, 3, n)
	})
	t.Run("base:[4, 3], min:[4, 3]", func(t *testing.T) {
		s.Pop()

		n, err := s.GetMin()
		assert.Nil(t, err)
		assert.Equal(t, 3, n)
	})
	t.Run("base:[4], min:[4]", func(t *testing.T) {
		s.Pop()

		n, err := s.GetMin()
		assert.Nil(t, err)
		assert.Equal(t, 4, n)
	})
	t.Run("base:[], min:[]", func(t *testing.T) {
		s.Pop()

		_, err := s.GetMin()
		assert.Equal(t, ErrEmptyStack, err)
	})
}
