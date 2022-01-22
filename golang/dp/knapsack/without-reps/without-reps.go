package withoutreps

/*
W - вместимость
w - вес
с - стоимость
*/
func knapsack(W int, w []int, c []int) int {
	dp := make([][]int, len(w)+1)

	for i := 0; i < len(dp); i++ { // w
		dp[i] = make([]int, W+1)
		for j := 0; j < len(dp[i]); j++ { // W
			if i == 0 || j == 0 {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j]
				if j >= w[i-1] {
					dp[i][j] = max(dp[i][j], dp[i-1][j-w[i-1]]+c[i-1])
				}
			}
		}
	}

	return dp[len(w)][W]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
