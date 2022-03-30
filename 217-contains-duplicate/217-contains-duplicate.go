func containsDuplicate(nums []int) bool {
    maps := make(map[int]bool, 0)

	result := false

	for i := 0; i < len(nums); i++ {
		if maps[nums[i]] {
			result = true
			break
		}

		maps[nums[i]] = true
	}

	return result
}