package median_of_two_sorted_arrays

import "math"

// @see https://www.geeksforgeeks.org/median-of-two-sorted-arrays-of-different-sizes/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	n := len(nums1)
	m := len(nums2)
	if n > m {
		return findMedianSortedArrays(nums2, nums1) // Swapping to make A smaller
	}
	start := 0
	end := n
	realmidinmergedarray := (n + m + 1) / 2
	var (
		mid, leftAsize, leftBsize, leftA, leftB, rightA, rightB int
	)

	for start <= end {
		mid = (start + end) / 2
		leftAsize = mid
		leftBsize = realmidinmergedarray - mid
		// checking overflow
		// of indices
		leftA, leftB = math.MinInt64, math.MinInt64
		rightA, rightB = math.MaxInt64, math.MaxInt64
		if leftAsize > 0 {
			leftA = nums1[leftAsize-1]
		}
		if leftBsize > 0 {
			leftB = nums2[leftBsize-1]
		}
		if leftAsize < n {
			rightA = nums1[leftAsize]
		}
		if leftBsize < m {
			rightB = nums2[leftBsize]
		}

		// if correct partition is done
		if leftA <= rightB && leftB <= rightA {
			if (m+n)%2 == 0 {
				return (math.Max(float64(leftA), float64(leftB)) + math.Min(float64(rightA), float64(rightB))) / 2.0
			}
			return math.Max(float64(leftA), float64(leftB))
		} else if leftA > rightB {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0.0
}
