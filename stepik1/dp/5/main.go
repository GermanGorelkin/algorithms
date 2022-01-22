/*
https://stepik.org/lesson/13262/step/4?unit=3447
Задача на программирование: лестница
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

	n := readInt(reader, buf)
	arr := readSliceInt(reader, buf, n)

	result := climbingStairs2(arr)
	fmt.Printf("%d\n", result)
}

func climbingStairs(c []int) int {
	if len(c) == 1 {
		return c[0]
	}

	dp := make([]int, len(c))
	dp[0] = c[0]
	dp[1] = max(c[0]+c[1], c[1])

	for i := 2; i < len(c); i++ {
		dp[i] = max(dp[i-2], dp[i-1]) + c[i]
	}

	return dp[len(dp)-1]
}

func climbingStairs2(c []int) int {
	var dp2, dp1, curr int

	for _, v := range c {
		curr = max(dp2, dp1) + v
		dp2 = dp1
		dp1 = curr
	}
	return curr
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
