func maxSubarraySumCircular(nums []int) int {
	
	size := len(nums)

	dp := make([]int, 0)
	dp = append(dp, 0, nums[0])

	if size == 1 {
		return nums[0]
	}

	sum := 0

	for i := 0; i < size; i++ {
		sum += nums[i]
	}

	maxValue := getMaxValue(nums)
	minValue := getMinValue(nums)

	if sum == minValue {
		return maxValue
	}

	return max(maxValue, sum-minValue)
}


func getMaxValue(nums []int) int {
	dp := make([]int, 0)
	dp = append(dp, 0, nums[0])

	maxValue := nums[0]

	for i := 2; i <= len(nums); i++ {
		value := max(dp[i-1]+nums[i-1], nums[i-1])
		dp = append(dp, value)
		maxValue = max(value, maxValue)
	}

	return maxValue
}

func getMinValue(nums []int) int {
	dp := make([]int, 0)
	dp = append(dp, 0, nums[0])

	minValue := nums[0]

	for i := 2; i <= len(nums); i++ {
		value := min(dp[i-1]+nums[i-1], nums[i-1])
		dp = append(dp, value)
		minValue = min(value, minValue)
	}

	return minValue
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
