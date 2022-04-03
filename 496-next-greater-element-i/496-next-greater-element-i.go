func nextGreaterElement(nums1 []int, nums2 []int) []int {
    result := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		points := false
		grateVal := -1
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				points = true
			}

			if points && nums1[i] < nums2[j] {
				grateVal = nums2[j]
				break
			}
		}
		result = append(result, grateVal)
	}

	return result
}