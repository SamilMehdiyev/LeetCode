package solutions

// ValidPalindrome function is a solution for the
// Problem #680 - Valid Palindrome II
// from leetcode.com
func ValidPalindrome(s string) bool {
	return checkValidPalindrome(s, 1)
}

func checkValidPalindrome(s string, attempt int) bool {
	start := 0
	end := len(s) - 1

	for start < end {
		if s[start] != s[end] {
			if attempt == 0 {
				return false
			}
			fromStartRemoved := checkValidPalindrome(s[:start]+s[start+1:], attempt-1)
			fromEndRemoved := checkValidPalindrome(s[:end]+s[end+1:], attempt-1)

			return fromStartRemoved || fromEndRemoved
		}
		start++
		end--
	}

	return true
}
