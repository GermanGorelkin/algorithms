/*
https://stepik.org/lesson/45970/step/3?unit=24123
Проверка более общего свойства дерева поиска
*/

package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_isValidBST(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `2 1 2
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		assert.True(t, isValidBST(root, nil, nil))
	})
	t.Run("2", func(t *testing.T) {
		s := `1 1 2
		2 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		assert.False(t, isValidBST(root, nil, nil))
	})
	t.Run("3", func(t *testing.T) {
		s := `2 1 2
		1 -1 -1
		2 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		assert.True(t, isValidBST(root, nil, nil))
	})
	t.Run("4", func(t *testing.T) {
		s := `2 1 2
		2 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		assert.False(t, isValidBST(root, nil, nil))
	})
	t.Run("1", func(t *testing.T) {
		s := `2147483647 -1 -1`
		r := strings.NewReader(s)
		n := 1

		root := treeBuild(n, r)

		assert.True(t, isValidBST(root, nil, nil))
	})
}
