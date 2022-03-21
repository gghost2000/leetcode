func lengthOfLIS(nums []int) int {
    n := len(nums)

	dp := make([]int, 0)

	for i := 0; i < n; i++ {
		dp = append(dp, 1)
	}

	maxV := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxV = max(maxV, dp[i])
	}

	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}