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
	realmidinmergedarray := (n + m + 1) >> 1 // /2
	var (
		mid, leftAsize, leftBsize, leftA, leftB, rightA, rightB int
	)
	odd := (m+n)%2 == 0
	for start <= end {
		mid = (start + end) >> 1 // /2
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
			if odd {
				return float64(Max(leftA, leftB)+Min(rightA, rightB)) / 2.0
			}
			return float64(Max(leftA, leftB))
		} else if leftA > rightB {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return 0.0
}

// Max returns the largest of x or y.
func Max(x, y int) int {
	if x < y {
		return y
	}
	return x
}

// Min returns the smallest of x or y.
func Min(x, y int) int {
	if x > y {
		return y
	}
	return x
}
