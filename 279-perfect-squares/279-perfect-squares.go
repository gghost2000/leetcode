func numSquares(n int) int {
    dp := make([]int, n+1)
	dp[0] = 0

	sq := make([]int, 0)

	for i := 1; i <= n; i++ {

		if i*i <= n {
			sq = append(sq, i*i)
		}

		dp[i] = 1000

		for j := 0; j < len(sq); j++ {
			if i-sq[j] >= 0 {
				dp[i] = min(dp[i], dp[i-sq[j]] + 1)
			}
		}
	}

	return dp[n]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
