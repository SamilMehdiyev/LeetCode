package solutions

// FindAnagrams function is a solution for the
// Problem #438 - Find All Anagrams in a String
// from leetcode.com
func FindAnagrams(s string, p string) []int {
	if s == "" || p == "" || len(s) < len(p) {
		return []int{}
	}

	idxs := []int{}
	chrs := make([]int, 26)

	for i := 0; i < len(p); i++ {
		chrs[p[i]-'a']++
	}

	start, end, length := 0, 0, len(p)
	diff := length

	for ; end < length; end++ {
		tmp := s[end]
		chrs[tmp-'a']--
		if chrs[tmp-'a'] >= 0 {
			diff--
		}
	}

	if diff == 0 {
		idxs = append(idxs, 0)
	}

	for end < len(s) {
		tmp := s[start]

		if chrs[tmp-'a'] >= 0 {
			diff++
		}

		chrs[tmp-'a']++
		start++

		tmp = s[end]
		chrs[tmp-'a']--

		if chrs[tmp-'a'] >= 0 {
			diff--
		}

		if diff == 0 {
			idxs = append(idxs, start)
		}

		end++
	}

	return idxs
}
