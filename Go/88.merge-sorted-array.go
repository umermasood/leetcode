package leetcode

import "sort"

func merge(nums1 []int, m int, nums2 []int, n int) {
	i := m
	for _, v := range nums2 {
		nums1[i] = v
		i += 1
	}
	sort.Ints(nums1)
}
