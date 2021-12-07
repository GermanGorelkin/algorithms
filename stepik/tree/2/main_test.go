package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_validatorV1(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `2 1 2
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		var vl validatorV1
		assert.True(t, vl.isValidBST(root))
	})
	t.Run("2", func(t *testing.T) {
		s := `1 1 2
		2 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		var vl validatorV1
		assert.False(t, vl.isValidBST(root))
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

		var vl validatorV1
		assert.True(t, vl.isValidBST(root))
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

		var vl validatorV1
		assert.True(t, vl.isValidBST(root))
	})
	t.Run("4", func(t *testing.T) {
		s := `4 1 -1
		2 2 3
		1 -1 -1
		5 -1 -1`
		r := strings.NewReader(s)
		n := 4

		root := treeBuild(n, r)

		var vl validatorV1
		assert.False(t, vl.isValidBST(root))
	})
}

func Test_validatorV2(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		s := `2 1 2
		1 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		var vl validatorV2
		assert.True(t, vl.isValidBST(root, nil, nil))
	})
	t.Run("2", func(t *testing.T) {
		s := `1 1 2
		2 -1 -1
		3 -1 -1`
		r := strings.NewReader(s)
		n := 3

		root := treeBuild(n, r)

		var vl validatorV2
		assert.False(t, vl.isValidBST(root, nil, nil))
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

		var vl validatorV2
		assert.True(t, vl.isValidBST(root, nil, nil))
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

		var vl validatorV2
		assert.True(t, vl.isValidBST(root, nil, nil))
	})
	t.Run("4", func(t *testing.T) {
		s := `4 1 -1
		2 2 3
		1 -1 -1
		5 -1 -1`
		r := strings.NewReader(s)
		n := 4

		root := treeBuild(n, r)

		var vl validatorV2
		assert.False(t, vl.isValidBST(root, nil, nil))
	})
}
