package with_reps

/*
W - вместимость
w - вес
с - стоимость
*/
func knapsack(W int, w []int, c []int) int {
	dp := make([]int, W+1)
	var result int

	for i := 0; i <= W; i++ {
		for j := 0; j < len(w); j++ {
			if i >= w[j] {
				dp[i] = max(dp[i], dp[i-w[j]]+c[j])
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
