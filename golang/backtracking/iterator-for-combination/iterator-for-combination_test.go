package iterator_for_combination

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CombinationIterator(t *testing.T) {
	t.Run("abc_2", func(t *testing.T) {
		iter := Constructor("abc", 2)

		assert.True(t, iter.HasNext())
		assert.Equal(t, "ab", iter.Next())

		assert.True(t, iter.HasNext())
		assert.Equal(t, "ac", iter.Next())

		assert.True(t, iter.HasNext())
		assert.Equal(t, "bc", iter.Next())

		assert.False(t, iter.HasNext())
	})
}
