func coinChange(coins []int, amount int) int {

	dp := make([]int, amount+1)

	for i := 1; i < len(dp); i++ {
		dp[i] = 5001
	}

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] = min(dp[j], dp[j-coins[i]]+1)
		}
	}

	result := dp[amount]

	if result == 5001 {
		result = -1
	}

	return result
}


func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}
