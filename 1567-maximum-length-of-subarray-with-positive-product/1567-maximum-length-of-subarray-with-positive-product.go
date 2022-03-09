func getMaxLen(nums []int) int {
    size := len(nums)
	maxLen := 0

	zero := make([]int, 0)

	start := 0

	for i := 0; i < size; i++ {
		if nums[i] == 0 {
			len := getMax(zero, start, i-1)
			zero = make([]int, 0)
			maxLen = max(maxLen, len)
			start = i + 1
		} else {
			if nums[i] < 0 {
				zero = append(zero, i)
			}
		}
	}

	maxLen = max(getMax(zero, start, size-1), maxLen)

	return maxLen

}

func getMax(zero []int, start int, end int) int {
	if start > end {
		return 0
	}

	size := len(zero)

	if size%2 == 0 {
		return end - start + 1
	} else {
		return max(end-zero[0], zero[size-1]-start)
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
