func minCostClimbingStairs(cost []int) int {
    dp := make([]int, 0)
	dp = append(dp, cost[0], cost[1])
    
    length := len(cost)

	for i := 2; i < length; i++ {
		dp = append(dp, min(dp[i-1], dp[i-2])+cost[i])
	}
    
    return min(dp[length-1], dp[length-2])
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}