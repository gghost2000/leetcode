func wiggleMaxLength(nums []int) int {
   	n := len(nums)

	dp := make([][]int, 0)

	for i := 0; i < n; i++ {
		dp = append(dp, []int{1, 1})
	}

	maxV := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {

			if nums[i]-nums[j] > 0 {
				dp[i][0] = max(dp[j][1]+1, dp[i][0])
			} else if nums[i] == nums[j] {
				continue
			} else {
				dp[i][1] = max(dp[j][0]+1, dp[i][1])
			}
		}
		maxV = max(maxV, max(dp[i][1], dp[i][0]))
	}

	return maxV
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}