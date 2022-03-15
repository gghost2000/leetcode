func generate(numRows int) [][]int {
    dp := make([][]int, 0)
	dp = append(dp, []int{1})

	for i := 1; i < numRows; i++ {
		arr := make([]int, 0)
		arr = append(arr, 1)

		for j := 1; j < i; j++ {
			arr = append(arr, dp[i-1][j-1]+dp[i-1][j])
		}
		arr = append(arr, 1)

		dp = append(dp, arr)
	}

	return dp
}