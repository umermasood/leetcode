package leetcode

import "sort"

func majorityElement(nums []int) int {
	// The intuition behind this approach is that if an element occurs more than n/2 times
	// in the array (where n is the size of the array), it will always occupy the middle
	// position when the array is sorted. Therefore, we can sort the array and return the
	// element at index n/2.
	sort.Ints(nums)
	return nums[len(nums)/2]
}
