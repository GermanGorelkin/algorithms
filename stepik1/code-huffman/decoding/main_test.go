package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_decode(t *testing.T) {
	t.Run("abacabad", func(t *testing.T) {
		code := "01001100100111"
		codeMap := map[string]string{
			"0":   "a",
			"10":  "b",
			"110": "c",
			"111": "d",
		}
		want := "abacabad"
		assert.Equal(t, want, decode(code, codeMap))
	})
	t.Run("a", func(t *testing.T) {
		code := "0"
		codeMap := map[string]string{
			"0": "a",
		}
		want := "a"
		assert.Equal(t, want, decode(code, codeMap))
	})
}
