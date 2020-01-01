package balanced_parentheses

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValid(t *testing.T) {
	t.Run("()", func(t *testing.T) {
		assert.True(t, IsValid("()"))
	})
	t.Run("()[]{}", func(t *testing.T) {
		assert.True(t, IsValid("()[]{}"))
	})
	t.Run("(]", func(t *testing.T) {
		assert.False(t, IsValid("(]"))
	})
	t.Run("([)]", func(t *testing.T) {
		assert.False(t, IsValid("([)]"))
	})
	t.Run("(", func(t *testing.T) {
		assert.False(t, IsValid("("))
	})
	t.Run("]", func(t *testing.T) {
		assert.False(t, IsValid("]"))
	})
}
