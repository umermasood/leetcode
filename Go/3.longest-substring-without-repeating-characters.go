package leetcode

import (
	"strings"
)

func lengthOfLongestSubstring(str string) int {
	longestSubstringLength := 0
	substringLength := 0

	for i := range str {
		curSubstr := ""
		for _, cci := range str[i:] {
			if !strings.Contains(curSubstr, string(cci)) {
				curSubstr += string(cci)
			} else {
				break
			}
			substringLength = len(curSubstr)
			if substringLength > longestSubstringLength {
				longestSubstringLength = substringLength
			}
		}
	}

	return longestSubstringLength
}
