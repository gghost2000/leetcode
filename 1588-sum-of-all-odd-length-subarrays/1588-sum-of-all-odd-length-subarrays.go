func sumOddLengthSubarrays(arr []int) int {
    result := 0

	for i := 1; i <= len(arr); i += 2 {
		for j := 0; j+i <= len(arr); j++ {
			result += getSum(arr[j : j+i])
		}
	}

	return result
}


func getSum(a []int) int {

	result := 0
	for i := 0; i < len(a); i++ {
		result += a[i]
	}

	return result
}