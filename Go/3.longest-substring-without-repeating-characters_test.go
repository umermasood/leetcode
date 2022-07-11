package leetcode

import "testing"

func Test_longest_substring_without_repeating_characters(t *testing.T) {

	if result := lengthOfLongestSubstring("abcabcbb"); result != 3 {
		t.Errorf("Error")
	}
	if result := lengthOfLongestSubstring(" "); result != 1 {
		t.Errorf("Error")
	}
	if result := lengthOfLongestSubstring("cbeabcde"); result != 5 {
		t.Errorf("Error")
	}
}
