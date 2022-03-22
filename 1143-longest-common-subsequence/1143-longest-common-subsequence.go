func longestCommonSubsequence(text1 string, text2 string) int {
 	n, m := len(text1), len(text2)

	dp := make([][]int, 0)

	for i := 0; i <= m; i++ {
		arr := make([]int, n+1)
		dp = append(dp, arr)
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if text1[j-1] == text2[i-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

    return dp[m][n]
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
