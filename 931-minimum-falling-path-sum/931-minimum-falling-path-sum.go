func minFallingPathSum(matrix [][]int) int {
    dp := make([][]int, 0)
	dp = append(dp, matrix[0])

	left := 0
	mid := 0
	right := 0

	total := 10000

	for i := 1; i < len(matrix); i++ {
		arr := make([]int, 0)
		for j := 0; j < len(matrix[0]); j++ {

			left = j - 1
			mid = j
			right = j + 1

			if left == -1 {
				left = 0
			}

			if right == len(matrix[0]) {
				right = len(matrix[0]) - 1
			}

			arr = append(arr, findMin(dp[i-1][left], findMin(dp[i-1][mid], dp[i-1][right]))+matrix[i][j])
		}

		dp = append(dp, arr)
	}

	for i := 0; i < len(matrix); i++ {
		total = findMin(total, dp[len(matrix)-1][i])
	}

	return total
}

func findMin(a, b int) int {
	if a > b {
		return b
	}

	return a
}
