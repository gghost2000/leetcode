func searchMatrix(matrix [][]int, target int) bool {
	arr := make([]int, 0)

	for i := 0; i < len(matrix); i++ {
		arr = append(arr, matrix[i]...)
	}

	low := 0
	high := len(arr) - 1

	for low <= high {
        mid := (low + high)/2

		if arr[mid] == target {
			return true 
		} else if arr[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}

	}

	return false
}