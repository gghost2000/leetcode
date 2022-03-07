func jump(nums []int) int {
    size := len(nums)

	dp := make([]int, 10000)
	dp[0] = 0

	for i := 1; i < size; i++ {
		dp[i] = 10000
	}

	for i := 0; i < size; i++ {
		reachValue := i + nums[i]

		for j := i + 1; j <= reachValue; j++ {

			if j >= size {
				break
			}

			dp[j] = min(dp[j], dp[i]+1)
		}
	}

	return dp[size-1]
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}