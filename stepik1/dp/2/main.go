/*
Задача на программирование: наибольшая невозрастающая подпоследовательность
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

	result := LDS(arr)
	fmt.Printf("%d\n", len(result))
	for _, v := range result {
		fmt.Printf("%d ", v)
	}
}

func ceilIndex(nums []int, increasingSub []int, end int, currIdx int) int {
	low := 1
	high := end
	for low <= high {
		mid := low + (high-low)/2
		if nums[increasingSub[mid]] >= nums[currIdx] {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}

func LDS(nums []int) []int {
	parent := make([]int, len(nums))
	increasingSub := make([]int, len(nums)+1)
	length := 0

	for i := 0; i < len(nums); i++ {
		low := ceilIndex(nums, increasingSub, length, i)
		pos := low
		parent[i] = increasingSub[pos-1]
		increasingSub[pos] = i

		if pos > length {
			length = pos
		}
	}

	LIS := make([]int, length)
	k := increasingSub[length]
	for j := length - 1; j >= 0; j-- {
		LIS[j] = k
		k = parent[k]
	}

	return LIS
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
