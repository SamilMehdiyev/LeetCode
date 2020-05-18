package solutions

// MaxPower function is a solution for the
// Problem #1446 - Consecutive Characters
// from leetcode.com
func MaxPower(s string) int {
	if len(s) == 0 {
		return 0
	}

	maxLength := 1
	length := 1

	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			length++
			if length > maxLength {
				maxLength = length
			}
		} else {
			length = 1
		}
	}

	if length > maxLength {
		maxLength = length
	}

	return maxLength
}
