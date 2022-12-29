package median_of_two_sorted_arrays

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	i := ((l1 + l2) >> 1)
	if (l1+l2)%2 == 1 {
		if i >= l1 {
			return float64(nums2[i-l1])
		} else {
			return float64(nums1[i])
		}
	} else {
		var a, b int
		if i-1 >= l1 {
			a = nums2[i-1-l1]
		} else {
			a = nums1[i-1]
		}
		if i >= l1 {
			b = nums2[i-l1]
		} else {
			b = nums1[i]
		}
		return (float64(a + b)) / 2
	}
	return 0
}
