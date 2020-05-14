package solutions

// FindTheDifference function is a solution for the
// Problem #389 - Find the Difference
// from leetcode.com
func FindTheDifference(s string, t string) byte {
	chars := make([]int, 26)

	for i := 0; i < len(s); i++ {
		chars[s[i]-'a']++
		chars[t[i]-'a']--
	}
	chars[t[len(t)-1]-'a']--

	for i := 0; i < len(chars); i++ {
		if chars[i] < 0 {
			return byte(i + 'a')
		}
	}
	return 'a'
}
