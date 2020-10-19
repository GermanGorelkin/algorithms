package main

import (
	"fmt"
)

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	data := make([]int, n)
	nums := make([]int, 11)
	var num int
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &num)
		data[i] = num
		nums[num]++
	}

	var indx int
	for j := 1; j < len(nums); j++ {
		for i := 0; i < nums[j]; i++ {
			data[indx] = j
			indx++
		}
	}

	for i := 0; i < len(data); i++ {
		fmt.Printf("%d ", data[i])
	}
}
