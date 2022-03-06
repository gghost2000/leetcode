func deleteAndEarn(nums []int) int {
    dp := make([]int, 0)

	countMap := make(map[int]int)

	maxNum := 0
	for _, num := range nums {
		countMap[num] += num
		maxNum = max(num, maxNum)
	}

	if len(countMap) == 1 {
		return countMap[nums[0]]
	}

	dp = append(dp, 0, countMap[1])

	for i := 2; i <= maxNum; i++ {
		dp = append(dp, max(dp[i-1], dp[i-2]+countMap[i]))
	}

	return dp[maxNum]
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}

