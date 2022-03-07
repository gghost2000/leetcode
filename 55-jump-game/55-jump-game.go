func canJump(nums []int) bool {
    maxNum := 0
    size := len(nums)
    
    if size == 1 {
		return true
	}
    
    for i := 1; i < size; i++ {
		maxNum = max(maxNum, i+nums[i-1])

        
        if maxNum <= i {
			return false
		}
	}
    
    if maxNum >= size {
		return true
	}
    
    return false
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
