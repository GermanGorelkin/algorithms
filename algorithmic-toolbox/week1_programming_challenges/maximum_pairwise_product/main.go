package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		_, _ = fmt.Scanf("%d", &numbers[i])
	}

	fmt.Println(maximum_pairwise_product(numbers))
}

func maximum_pairwise_product(numbers []int) int {
	var max1, max2 int
	if numbers[0] > numbers[1] {
		max1, max2 = numbers[0], numbers[1]
	} else {
		max1, max2 = numbers[1], numbers[0]
	}

	for i := 2; i < len(numbers); i++ {
		if numbers[i] > max1 {
			max2 = max1
			max1 = numbers[i]
		} else if numbers[i] > max2 {
			max2 = numbers[i]
		}
	}
	return max1 * max2
}
