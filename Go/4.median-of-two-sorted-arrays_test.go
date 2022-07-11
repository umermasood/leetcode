package leetcode

import "testing"

func Test_median_of_two_sorted_arrays(t *testing.T) {

	if result := findMedianSortedArrays([]int{1, 3}, []int{2}); result != 2.0 {
		t.Errorf("Error")
	}
	if result := findMedianSortedArrays([]int{1, 2}, []int{3, 4}); result != 2.5 {
		t.Errorf("Error")
	}
}
