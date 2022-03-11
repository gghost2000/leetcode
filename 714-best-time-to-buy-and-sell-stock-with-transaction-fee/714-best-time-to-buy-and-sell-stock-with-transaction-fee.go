func maxProfit(prices []int, fee int) int {
    size := len(prices)

	dp := make([][2]int, 50000)

	dp[0][0] = -prices[0]
	dp[0][1] = 0         

	for i := 1; i < size; i++ {
		dp[i][0] = max(dp[i-1][1]-prices[i], dp[i-1][0])     
		dp[i][1] = max(dp[i-1][0]+prices[i]-fee, dp[i-1][1]) 
	}

	return max(dp[size-1][0], dp[size-1][1])
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}