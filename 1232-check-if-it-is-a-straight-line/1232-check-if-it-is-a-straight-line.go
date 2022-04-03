func checkStraightLine(coordinates [][]int) bool {
    
	sort.Slice(coordinates, func(i, j int) bool {
		if coordinates[i][0] < coordinates[j][0] {
			return true
		} else if coordinates[i][0] == coordinates[j][0] {
			if coordinates[i][1] < coordinates[j][1] {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	})

	dx := coordinates[1][0] - coordinates[0][0]
	dy := coordinates[1][1] - coordinates[0][1]

	result := true

	for i := 2; i < len(coordinates); i++ {

		if dx*(coordinates[i][1]-coordinates[i-1][1]) != dy*(coordinates[i][0]-coordinates[i-1][0]) {
			result = false
			break
		}

	}

	return result
}