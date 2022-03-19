func maximalSquare(matrix [][]byte) int {
	m, n := len(matrix), len(matrix[0])

	dp := make([][]int, 0)

	total := getInt(matrix[0][0])
	for i := 0; i < m; i++ {
		arr := make([]int, n)

		dp = append(dp, arr)
	}

	dp[0][0] = getInt(matrix[0][0])

	for i := 1; i < n; i++ {
		dp[0][i] = getInt(matrix[0][i])
		total = max(total, dp[0][i])
	}

	for i := 1; i < m; i++ {
		dp[i][0] = getInt(matrix[i][0])
		total = max(total, dp[i][0])
	}


	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
            if getInt(matrix[i][j]) == 0 {
				dp[i][j] = 0
				continue
			}
            
			a1 := dp[i-1][j-1]
			a2 := dp[i-1][j]
			a3 := dp[i][j-1]

			dp[i][j] = min(a1, min(a2, a3)) + getInt(matrix[i][j])
			total = max(total, dp[i][j])
		}
	}

	return total * total
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func getInt(b byte) int {
	if b == '0' {
		return 0
	} else {
		return 1
	}
}
