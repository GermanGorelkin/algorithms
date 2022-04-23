package main

import (
	"testing"
)

func BenchmarkSolv(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Solv(100_000, 2)
	}
}
