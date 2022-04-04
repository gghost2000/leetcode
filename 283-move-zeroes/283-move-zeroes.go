func moveZeroes(nums []int)  {
    
	index := 0
	count := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			count++
		} else {
			nums[index] = nums[i]
			index++
		}
	}

	for i := 0; i < count; i++ {
		nums[index+i] = 0
	}
}