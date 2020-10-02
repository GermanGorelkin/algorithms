package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_bsearch(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		data := []int{0, 1, 2, 3, 4, 5, 6}

		assert.Equal(t, 1, bsearch(data, 0))
		assert.Equal(t, 2, bsearch(data, 1))
		assert.Equal(t, 4, bsearch(data, 3))
		assert.Equal(t, 5, bsearch(data, 4))
		assert.Equal(t, 6, bsearch(data, 5))
		assert.Equal(t, 7, bsearch(data, 6))
	})
}
