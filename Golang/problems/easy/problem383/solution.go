package solutions

// CanConstruct function is a solution for the
// Problem #383 - Ransom Note
// from leetcode.com
func CanConstruct(ransomNote string, magazine string) bool {

	magazineMap := make(map[rune]int)

	for _, ch := range magazine {
		magazineMap[ch]++
	}

	for _, ch := range ransomNote {
		if _, found := magazineMap[ch]; found {
			if magazineMap[ch] > 0 {
				magazineMap[ch]--
			} else {
				return false
			}
		} else {
			return false
		}
	}

	return true
}
