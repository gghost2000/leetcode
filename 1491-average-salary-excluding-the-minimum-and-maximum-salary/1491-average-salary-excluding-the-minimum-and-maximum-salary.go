func average(salary []int) float64 {
    min := salary[0]
	max := salary[0]

	sum := 0.0
	for i := 0; i < len(salary); i++ {
		sum += float64(salary[i])
		max = maxV(max, salary[i])
		min = minV(min, salary[i])
	}

	result := float64((sum - float64(max) - float64(min)) / float64(len(salary)-2))
    
    return result
}


func maxV(a, b int) int {
	if a > b {
		return a
	}

	return b
}

func minV(a, b int) int {
	if a > b {
		return b
	}
	return a
}
