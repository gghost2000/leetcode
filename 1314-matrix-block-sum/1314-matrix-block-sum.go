func matrixBlockSum(mat [][]int, k int) [][]int {
output := make([][]int, 0)

	width := len(mat[0])
	length := len(mat)

	for i := 0; i < length; i++ {
		arr := make([]int, 0)

		for j := 0; j < width; j++ {
			val := 0
			for r := i - k; r <= i+k; r++ {

				if r < 0 || r >= length {
					continue
				}

				for c := j - k; c <= j+k; c++ {

					if c < 0 || c >= width {
						continue
					}

					val += mat[r][c]

				}
			}
			arr = append(arr, val)
		}

		output = append(output, arr)
	}

	return output
}