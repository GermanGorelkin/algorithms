package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var count int
	_, _ = fmt.Scanf("%d", &count)
	reader := bufio.NewReader(os.Stdin)
	line2, _ := reader.ReadString('\n')
	line2 = strings.Trim(line2, "\n")

	var ss []string
	ss = strings.Split(line2, " ")
	data := make([]int, 0, len(ss))
	for i := 0; i < len(ss); i++ {
		v, _ := strconv.Atoi(ss[i])
		data = append(data, v)
	}

	_, cntInv := countInv(data)

	fmt.Printf("%d ", cntInv)
}

func countInv(data []int) ([]int, int) {
	l := len(data)
	if l == 1 {
		return data, 0
	}
	left, leftInv := countInv(data[:l/2])
	right, rightInv := countInv(data[l/2:])
	res, cntInv := merge(left, right)

	return res, leftInv + rightInv + cntInv
}

func merge(left, right []int) ([]int, int) {
	lsize, rsize := len(left), len(right)
	res := make([]int, lsize+rsize)
	var cntInt int
	l, r := 0, 0

	for i := 0; i < lsize+rsize; i++ {
		if l >= lsize {
			copy(res[i:], right[r:])
			cntInt += lsize - l
			break
		}
		if r >= rsize {
			copy(res[i:], left[l:])
			break
		}

		if left[l] <= right[r] {
			res[i] = left[l]
			l++
		} else {
			res[i] = right[r]
			r++
			cntInt += lsize - l
		}
	}

	return res, cntInt
}
