func twoSum(nums []int, target int) []int {
    maps := make(map[int]int, 0)

	result := make([]int, 0)

	for i := 0; i < len(nums); i++ {
		needVal := target - nums[i]

		if maps[needVal] != 0 {
			result = append(result, maps[needVal]-1, i)
			break
		} else {
			maps[nums[i]] = i + 1
		}
	}

	return result
}