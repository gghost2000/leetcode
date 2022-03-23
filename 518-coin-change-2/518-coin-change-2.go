func change(amount int, coins []int) int {
	dp := make([]int, amount+1)

	dp[0] = 1

	sort.Slice(coins, func(i, j int) bool {
		return coins[i] < coins[j]
	})

	for i := 0; i < len(coins); i++ {
		for j := coins[i]; j <= amount; j++ {
			dp[j] += dp[j-coins[i]]
		}
	}

	result := dp[amount]

	return result
}