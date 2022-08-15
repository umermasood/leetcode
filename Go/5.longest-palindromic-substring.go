package leetcode

func isPalindrome(sub string, i int) bool {
	if i > len(sub)/2 {
		return true
	}
	return sub[i] == sub[len(sub)-i-1] && isPalindrome(sub, i+1)
}

func longestPalindrome(s string) string {
	var longestPalindromicSubstring string
	for i, os := range s {
		curSub := string(os)
		for _, is := range s[i+1:] {
			curSub += string(is)
			if isPalindrome(curSub, 0) && len(curSub) > len(longestPalindromicSubstring) {
				longestPalindromicSubstring = curSub
			}
		}
	}
	return longestPalindromicSubstring
}
