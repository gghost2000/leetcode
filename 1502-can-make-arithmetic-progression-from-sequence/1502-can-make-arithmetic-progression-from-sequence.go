func canMakeArithmeticProgression(arr []int) bool {
    sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	gap := arr[1] - arr[0]

	result := true

	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != gap {
			result = false
			break
		}
	}

	return result
}