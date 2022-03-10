func maxProfit(prices []int) int {
    dp := make([]int, 0)
	dp = append(dp, 10001)

	size := len(prices)
	maxVal := 0

	for i := 0; i < size; i++ {
		dp = append(dp, min(dp[i], prices[i]))
		if prices[i] > dp[i+1] {
			maxVal = max(maxVal, prices[i]-dp[i+1])
		}
	}

	return maxVal
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}