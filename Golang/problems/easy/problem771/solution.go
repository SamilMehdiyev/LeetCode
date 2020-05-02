package solutions

// NumJewelsInStones function is a solution for the
// Problem #771 - Jewels and Stones
// from leetcode.com
func NumJewelsInStones(J string, S string) int {

	dictionary := make(map[rune]int)

	for _, ch := range S {
		dictionary[ch] = dictionary[ch] + 1
	}

	count := 0

	for _, ch := range J {
		count += dictionary[ch]
	}

	return count
}
