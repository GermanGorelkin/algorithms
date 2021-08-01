package main

import "fmt"

func main() {
	var n int
	_, _ = fmt.Scanf("%d", &n)

	ans := calc(n)

	fmt.Printf("%d\n", len(ans)-1)

	for _, v := range ans {
		fmt.Printf("%d ", v)
	}
}

func calc(n int) []int {
	dp := make([]int, n+1)
	dp[1] = 1

	for i := 2; i < len(dp); i++ {
		dp[i] = dp[i-1] + 1
		if i%2 == 0 {
			dp[i] = min(dp[i], dp[i/2]+1)
		}
		if i%3 == 0 {
			dp[i] = min(dp[i], dp[i/3]+1)
		}
	}

	k := dp[len(dp)-1]
	ans := make([]int, k)
	ans[k-1] = n

	for i := n; i > 1; {
		k--
		if i%3 == 0 && dp[i/3] == k {
			i /= 3
		} else if i%2 == 0 && dp[i/2] == k {
			i /= 2
		} else {
			i--
		}
		ans[k-1] = i
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
