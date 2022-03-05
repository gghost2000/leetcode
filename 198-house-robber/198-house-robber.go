func rob(nums []int) int {
    n := len(nums)
	dp := make([][2]int, 0)
	dp = append(dp, [2]int{0, 0}, [2]int{nums[0], 0})

	i := 2
	for i <= n {
		dp = append(dp,[2]int{findMax(dp[i-2][0], dp[i-2][1]) + nums[i-1], findMax(dp[i-1][0], dp[i-1][1])})
		i++
	}
    
    return findMax(dp[n][0], dp[n][1])

}



func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}