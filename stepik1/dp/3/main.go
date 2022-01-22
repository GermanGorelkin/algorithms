/*
Задача на программирование: расстояние редактирования
https://stepik.org/lesson/13258/step/8?unit=3443
*/

package main

import "fmt"

func main() {
	var s1, s2 string
	_, _ = fmt.Scanln(&s1)
	_, _ = fmt.Scanln(&s2)

	ans := EditDistance(s1, s2)
	fmt.Printf("%d", ans)
}

func EditDistance(s1, s2 string) int {
	n, m := len(s1)+1, len(s2)+1
	dp := make([][]int, n)

	for i := 0; i < n; i++ {
		dp[i] = make([]int, m)
		for j := 0; j < m; j++ {
			if i == 0 {
				dp[i][j] = j
			} else if j == 0 {
				dp[i][j] = i
			} else {
				c := diff(s1[i-1], s2[j-1])
				dp[i][j] = min3(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+c)
			}
		}
	}

	return dp[n-1][m-1]
}

func min3(a, b, c int) int {
	if a < b {
		return min2(a, c)
	}
	return min2(b, c)
}

func min2(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func diff(a, b byte) int {
	if a == b {
		return 0
	}
	return 1
}
