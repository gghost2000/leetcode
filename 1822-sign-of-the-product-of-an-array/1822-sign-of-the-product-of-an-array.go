func arraySign(nums []int) int {
  	minusCount := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {		
			return 0
		}

		if nums[i] < 0 {
			minusCount++
		}
	}

	if minusCount%2 == 0 {
		return 1
	}

	return -1 
}