func rob(nums []int) int {
	dp := make([]int, 0)
	dp = append(dp, 0, nums[0])


	n := len(nums)

	if n == 1 {
		return nums[0]
	}

	for i := 2; i <= n; i++ {
		dp = append(dp, findMax(dp[i-1], dp[i-2]+nums[i-1]))
	}

	return findMax(dp[n], dp[n-1])

}



func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}