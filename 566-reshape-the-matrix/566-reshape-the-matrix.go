func matrixReshape(mat [][]int, r int, c int) [][]int {
    m, n := len(mat), len(mat[0])

	nums := make([]int, 0)
	result := make([][]int, 0)

	for i := 0; i < r; i++ {
		arr := make([]int, c)
		result = append(result, arr)
	}

	if r*c != m*n {
		return mat
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			nums = append(nums, mat[i][j])
		}
	}

	for i := 0; i < len(nums); i++ {
		result[i/c][i%c] = nums[i]
	}

	return result
}