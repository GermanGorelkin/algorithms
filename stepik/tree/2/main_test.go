package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validator_useInOrder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `2 1 2
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		var vl validator
		assert.True(t, vl.useInOrder(root))
	})
	t.Run("2", func(t *testing.T) {
		s := `1 1 2
		2 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		var vl validator
		assert.False(t, vl.useInOrder(root))
	})
	t.Run("3", func(t *testing.T) {
		s := `1 -1 1
		2 -1 2
		3 -1 3
		4 -1 4
		5 -1 -1`
		r := strings.NewReader(s)
		n := 5

		root := treeBuild(n, r)

		var vl validator
		assert.True(t, vl.useInOrder(root))
	})
	t.Run("4", func(t *testing.T) {
		s := `4 1 2
		2 3 4
		6 5 6
		1 -1 -1
		3 -1 -1
		5 -1 -1
		7 -1 -1`
		r := strings.NewReader(s)
		n := 7

		root := treeBuild(n, r)

		var vl validator
		assert.True(t, vl.useInOrder(root))
	})
	t.Run("4", func(t *testing.T) {
		s := `4 1 -1
		2 2 3
		1 -1 -1
		5 -1 -1`
		r := strings.NewReader(s)
		n := 4

		root := treeBuild(n, r)

		var vl validator
		assert.False(t, vl.useInOrder(root))
	})
}
