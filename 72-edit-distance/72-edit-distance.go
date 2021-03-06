func minDistance(word1 string, word2 string) int {
    n, m := len(word1), len(word2)

	dp := make([][]int, 0)

    for i := 0; i <= m; i++ {
		arr := make([]int, n+1)
		dp = append(dp, arr)
	}

	for i := 0; i <= m; i++ {
		dp[i][0] = i
	}

	for i := 0; i <= n; i++ {
		dp[0][i] = i
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if word1[j-1] == word2[i-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}

	return dp[m][n]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
