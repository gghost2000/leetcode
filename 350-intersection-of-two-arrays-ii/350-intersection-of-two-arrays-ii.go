func intersect(nums1 []int, nums2 []int) []int {
    maps := make(map[int]int, 0)

	result := make([]int, 0)

	for i := 0; i < len(nums1); i++ {
		maps[nums1[i]] += 1
	}

	for i := 0; i < len(nums2); i++ {
		if maps[nums2[i]] > 0 {
			result = append(result, nums2[i])
			maps[nums2[i]] -= 1
		}
	}

	return result
}