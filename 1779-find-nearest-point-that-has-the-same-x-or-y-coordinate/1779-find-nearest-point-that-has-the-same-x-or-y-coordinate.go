func nearestValidPoint(x int, y int, points [][]int) int {
    result := 10001
	index := -1
	for i := 0; i < len(points); i++ {
		if x == points[i][0] && abs(y, points[i][1]) < result {
			index = i
			result = abs(y, points[i][1])
		}

		if y == points[i][1] && abs(x, points[i][0]) < result {
			index = i
			result = abs(x, points[i][0])
		}
	}

	return index
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func abs(a, b int) int {
	if a-b > 0 {
		return a - b
	}

	return b - a
}