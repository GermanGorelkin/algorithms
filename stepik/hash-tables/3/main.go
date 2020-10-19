package main

import (
	"fmt"
)

func main() {
	var text string
	var pattern string

	_, _ = fmt.Scanf("%s", &pattern)
	_, _ = fmt.Scanf("%s", &text)

	lenPattern := len(pattern)
	degrees := calcDegrees(X, len(pattern))
	hPattern := getHash([]byte(pattern), degrees)

	hashChunk := getHash([]byte(text[:len(pattern)]), degrees)

	var found []int
	if hashChunk == hPattern && text[:lenPattern]==pattern {
		found = append(found, 0)
	}

	for i := 0; i < len(text)-len(pattern); i++ {
		firstCh := int(text[i])
		nextCh := int(text[i+lenPattern])

		hashChunk = ((hashChunk + P - (firstCh*degrees[lenPattern-1])%P) * X) % P
		hashChunk += nextCh
		hashChunk %= P

		if hashChunk == hPattern && text[i+1:i+lenPattern+1] == pattern{
			found = append(found, i + 1)
		}
	}

	for _, v := range found{
		fmt.Printf("%d ", v)
	}
}

const(
	X = 263
	P =  1_000_000_007
)

func getHash(b []byte, degreesX []int) int {
	var hash int
	for i := 0; i < len(b); i++ {
		hash += int(b[i]) * degreesX[len(b)-i-1]
		hash %= P
	}
	return hash
}

func calcDegrees(x int ,n int) []int {
	degrees := make([]int, n)
	degrees[0] = 1

	for i := 1; i < n; i++ {
		degrees[i] = degrees[i-1] * x % P
	}

	return degrees
}
