package main

import "fmt"

func main() {
	var N, M int
	fmt.Scanf("%d %d ", &N, &M)

	matrix := make([][]int, N)
	for i := 0; i < N; i++ {
		matrix[i] = make([]int, N)
	}

	for i := 0; i < M; i++ {
		var from, to int
		fmt.Scanf("%d %d ", &from, &to)
		matrix[from-1][to-1] = 1
	}

	var result bool
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if (matrix[i][j] == 0 && matrix[j][i] == 0) || (matrix[i][j] == 1 && matrix[j][i] == 1) {
				result = true
				break
			}
		}
	}

	if result {
		fmt.Println("NO ")
	} else {
		fmt.Println("YES ")
	}
}
