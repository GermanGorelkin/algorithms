package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHashTable(t *testing.T){
	t.Run("1", func(t *testing.T) {
		ht := NewHashTable()
		ht.add(911, "police")
		ht.add(76213, "Mom")
		ht.add(17239, "Bob")

		assert.Equal(t, "Mom", ht.find(76213))
		assert.Equal(t, "not found", ht.find(910))
		assert.Equal(t, "police", ht.find(911))

		ht.del(910)
		ht.del(911)

		assert.Equal(t, "not found", ht.find(911))
		assert.Equal(t, "Mom", ht.find(76213))

		ht.add(76213, "daddy")

		assert.Equal(t, "daddy", ht.find(76213))

	})
}
