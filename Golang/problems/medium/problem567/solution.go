package solutions

// CheckInclusion function is a solution for the
// Problem #567 - Permutation in String
// from leetcode.com
func CheckInclusion(s1 string, s2 string) bool {
	if s1 == "" || s2 == "" || len(s2) < len(s1) {
		return false
	}

	chrs := make([]int, 26)

	for i := 0; i < len(s1); i++ {
		chrs[s1[i]-'a']++
	}

	start, end, length := 0, 0, len(s1)
	diff := length

	for ; end < length; end++ {
		tmp := s2[end]
		chrs[tmp-'a']--
		if chrs[tmp-'a'] >= 0 {
			diff--
		}
	}

	if diff == 0 {
		return true
	}

	for end < len(s2) {
		tmp := s2[start]

		if chrs[tmp-'a'] >= 0 {
			diff++
		}

		chrs[tmp-'a']++
		start++

		tmp = s2[end]
		chrs[tmp-'a']--

		if chrs[tmp-'a'] >= 0 {
			diff--
		}

		if diff == 0 {
			return true
		}

		end++
	}

	return false
}
