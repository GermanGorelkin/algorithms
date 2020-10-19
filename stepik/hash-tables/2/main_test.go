package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_getHash(t *testing.T) {
	calcDegreeX()

	t.Run("world", func(t *testing.T) {
		s := "world"
		got := getHash(s, 5)
		want := 4
		assert.Equal(t, want, got)
	})

	t.Run("HellO", func(t *testing.T) {
		s := "HellO"
		got := getHash(s, 5)
		want := 4
		assert.Equal(t, want, got)
	})

	t.Run("luck", func(t *testing.T) {
		s := "luck"
		got := getHash(s, 5)
		want := 2
		assert.Equal(t, want, got)
	})

	t.Run("GooD", func(t *testing.T) {
		s := "GooD"
		got := getHash(s, 5)
		want := 2
		assert.Equal(t, want, got)
	})

	//t.Run("aaaaaaaaaaaaaaa", func(t *testing.T) {
	//	s := "aaaaaaaaaaaaaaa"
	//	got := getHash(s, 5)
	//	want := 2
	//	assert.Equal(t, want, got)
	//})
}