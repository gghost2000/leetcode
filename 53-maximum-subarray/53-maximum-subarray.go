func maxSubArray(nums []int) int {
    size := len(nums)
    
	if size == 1 {
		return nums[0]
	}

    
	dp := make([]int, 0)
	dp = append(dp, 0, nums[0])

	maxValue := nums[0]

	for i := 2; i <= size; i++ {
		value := max(dp[i-1]+nums[i-1], nums[i-1])
		dp = append(dp, value)
		maxValue = max(value, maxValue)
	}

	return maxValue
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}