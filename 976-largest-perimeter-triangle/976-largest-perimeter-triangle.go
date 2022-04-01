func largestPerimeter(nums []int) int {
    
	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	for i := len(nums) - 1; i >=2; i-- {
		if nums[i-2]+nums[i-1] > nums[i] {
			return nums[i-2] + nums[i-1] + nums[i]
		}
	}

	return 0
}