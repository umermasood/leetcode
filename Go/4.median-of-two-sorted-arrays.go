package leetcode

import "sort"

func findMedianSortedArrays(nums1, nums2 []int) float64 {
	merged := append(nums1, nums2...)
	sort.Ints(merged)

	length := len(merged)
	if length%2 == 0 {
		return float64(merged[length/2-1]+merged[length/2]) / 2
	}

	return float64(merged[length/2])
}
