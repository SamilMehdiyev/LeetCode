package solutions

// FirstUniqChar function is a solution for the
// Problem #387 - First Unique Character in a String
// from leetcode.com
func FirstUniqChar(s string) int {

	letterMap := make(map[rune]int)

	index := -1
	for _, ch := range s {
		letterMap[ch]++
	}

	for k, ch := range s {
		if letterMap[ch] == 1 {
			return k
		}
	}

	return index
}
