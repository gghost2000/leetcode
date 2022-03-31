func merge(nums1 []int, m int, nums2 []int, n int)  {
  for i := m; i < len(nums1); i++ {
		nums1[i] = nums2[i-m]
	}

	sort.Slice(nums1, func(i, j int) bool {
		return nums1[i] < nums1[j]
	})
	
}