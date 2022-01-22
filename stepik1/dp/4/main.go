/*
https://stepik.org/lesson/13259/step/5?unit=3444
Задача на программирование: рюкзак
*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024)
	var buf []byte

	W := readInt(reader, buf)
	n := readInt(reader, buf)
	w := readSliceInt(reader, buf, n)

	ans := knapsack(W, w)
	fmt.Printf("%d", ans)
}

func knapsack(W int, w []int) int {
	dp := make([][]int, len(w)+1)

	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, W+1)
		for j := 0; j < len(dp[i]); j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j]
				if j >= w[i-1] {
					dp[i][j] = max(dp[i][j], dp[i-1][j-w[i-1]]+w[i-1])
				}
			}
		}
	}

	return dp[len(w)][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

//---
func readSliceInt(r *bufio.Reader, buf []byte, size int) []int {
	s := make([]int, size)
	for i := range s {
		s[i] = readInt(r, buf)
	}
	return s
}

func readInt(r *bufio.Reader, buf []byte) (num int) {
	var err error
	var b byte
	buf = buf[:0]
	for {
		b, err = r.ReadByte()
		if b == ' ' || b == '\n' || err != nil {
			return parseInt(buf)
		} else {
			buf = append(buf, b)
		}
	}
}

func parseInt(b []byte) int {
	neg := false
	if b[0] == '-' {
		neg = true
		b = b[1:]
	}

	var n int
	for _, c := range b {
		c = c - '0'
		n *= 10
		n += int(c)
	}

	if neg {
		n = -n
	}
	return n
}
