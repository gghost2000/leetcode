func numberOfArithmeticSlices(nums []int) int {
    size := len(nums)

	if size <= 2 {
		return 0
	}

	l := 0
	r := 2
	count := 0

	for l < size-1 {
		if r >= size {
			l++
			r = l + 2
			continue
		}

		p := check(l, r, nums)

		if p {
			count++
			r++
		} else {
			l++
			r = l + 2
		}
	}

	return count
}

func check(l, r int, nums []int) bool {
	gap := nums[l+1] - nums[l]

	if r-l < 2 {
		return false
	}

	for i := l + 1; i <= r; i++ {
		if (nums[i] - nums[i-1]) != gap {
			return false
		}
	}

	return true
}
