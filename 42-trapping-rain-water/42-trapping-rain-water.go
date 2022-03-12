func trap(height []int) int {
    maxL := 0
	maxR := 0
	lP := 0
	rP := len(height) - 1
	total := 0

	for lP < rP {
		if height[lP] < height[rP] {
			if maxL <= height[lP] {
				maxL = height[lP]
			} else {
				total += maxL - height[lP]
			}
			lP++
		} else {
			if maxR <= height[rP] {
				maxR = height[rP]
			} else {
				total += maxR - height[rP]
			}
			rP--
		}
	}

	return total
}