func rob(nums []int) int {
	n := len(nums)
    
    if n== 1 {
		return nums[0]
	}
    
    return findMax(robMoney(1, n-1, nums), robMoney(2, n, nums))
}


func robMoney(start int, end int, nums []int) int {
	money := nums[start-1 : end]

	dp := make([]int, 0)
	dp = append(dp, 0, money[0])

	n := len(money)

	if start == end {
		return money[0]
	}

	for i := 2; i <= n; i++ {
		dp = append(dp, findMax(dp[i-1], dp[i-2]+money[i-1]))
	}

	return findMax(dp[n], dp[n-1])
}


func findMax(a, b int) int {
	if a > b {
		return a
	}

	return b
}