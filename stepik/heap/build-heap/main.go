package main

import (
	"fmt"
)

type step struct {
	i, j int
}

var steps []step

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &numbers[i])
	}

	buildHeap(numbers)

	fmt.Println(len(steps))
	for _, v := range steps {
		fmt.Printf("%d %d\n", v.i, v.j)
	}
}

func buildHeap(data []int) {
	for i := len(data)/2 - 1; i >= 0; i-- {
		siftDown(data, i)
	}
}

func siftDown(data []int, i int) {
	l := len(data) - 1

	for i*2+1 <= l {
		left := i*2 + 1
		right := i*2 + 2

		smallest := left

		if right <= l && data[right] < data[left] {
			smallest = right
		}

		if data[i] < data[smallest] {
			return
		}

		steps = append(steps, step{i: i, j: smallest})
		data[i], data[smallest] = data[smallest], data[i]
		i = smallest
	}
}
