package v2

/*
The point with in-degree - out-degree = N - 1 become the judge.
*/

func findJudge(N int, trust [][]int) int {
	if N == 1 && len(trust) == 0 {
		return 1
	}
	if N == 1 && len(trust) > 0 {
		return -1
	}

	count := make([]int, N+1)

	for i := 0; i < len(trust); i++ {
		count[trust[i][0]]--
		count[trust[i][1]]++
	}

	for i := 1; i < len(count); i++ {
		if count[i] == N-1 {
			return i
		}
	}

	return -1
}
