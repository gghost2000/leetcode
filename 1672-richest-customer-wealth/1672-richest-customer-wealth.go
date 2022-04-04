func maximumWealth(accounts [][]int) int {
    result := 0

	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		result = max(result, sum)
	}

	return result
}


func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
