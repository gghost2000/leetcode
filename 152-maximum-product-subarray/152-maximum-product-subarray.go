func maxProduct(nums []int) int {
    size := len(nums)

	maxDp := make([]int, 0)
	minDp := make([]int, 0)

	maxDp = append(maxDp, 0, nums[0])
	minDp = append(minDp, 0, nums[0])

	if size == 1 {
		return nums[0]
	}
    
    
	totalMaxValue := nums[0]
	for i := 2; i <= size; i++ {
		pMin := nums[i-1] * minDp[i-1]
		pMax := nums[i-1] * maxDp[i-1]

		maxValue := max(nums[i-1], max(pMin, pMax))
		minValue := min(nums[i-1], min(pMin, pMax))

		minDp = append(minDp, minValue)
		maxDp = append(maxDp, maxValue)

		totalMaxValue = max(maxValue, totalMaxValue)
	}

	return totalMaxValue
}

func min(a, b int) int {
	if a > b {
		return b
	}

	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}