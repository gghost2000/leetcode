func minPathSum(grid [][]int) int {
    m, n := len(grid), len(grid[0])

	dp := make([][]int, 0)

	for i := 0; i < m; i++ {
		arr := make([]int, n)

		dp = append(dp, arr)
	}

	dp[0][0] = grid[0][0]

	for i := 1; i < n; i++ {
		dp[0][i] = grid[0][i] + dp[0][i-1]
	}

	for i := 1; i < m; i++ {
		dp[i][0] = grid[i][0] + dp[i-1][0]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[i][j] = findMin(dp[i-1][j], dp[i][j-1]) + grid[i][j]
		}
	}

	return dp[m-1][n-1]
}


func findMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}