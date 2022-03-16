func minimumTotal(triangle [][]int) int {
    dp := make([][]int, 0)
	dp = append(dp, triangle[0])

	for i := 1; i < len(triangle); i++ {
		arr := make([]int, 0)

		for j := 0; j < len(triangle[i]); j++ {
			left := j - 1
			right := j

			if left == -1 {
				left = 0
			}

			if right >= len(triangle[i-1]) {
				right = len(triangle[i-1]) - 1
			}

			arr = append(arr, findMin(dp[i-1][left], dp[i-1][right])+triangle[i][j])

		}

		dp = append(dp, arr)
	}

	total := dp[len(triangle)-1][0]

	for i := 1; i < len(dp[len(triangle)-1]); i++ {
		total = findMin(total, dp[len(triangle)-1][i])
	}

	return total
}

func findMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
