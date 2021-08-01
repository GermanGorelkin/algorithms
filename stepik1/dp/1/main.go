/*
Задача на программирование: наибольшая последовательнократная подпоследовательность
https://stepik.org/lesson/13257/step/5?unit=3442
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

	result := ls(arr)
	fmt.Printf("%d\n", result)
}

func ls(arr []int) int {
	dp := make([]int, len(arr))
	var result int

	for i := 0; i < len(arr); i++ {
		dp[i] = 1
		for j := 0; j < i; j++ {
			if arr[i]%arr[j] == 0 {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		result = max(result, dp[i])
	}

	return result
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
