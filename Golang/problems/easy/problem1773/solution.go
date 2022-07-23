package solutions

// CountMatches function is a solution for the
// Problem #1773. Count Items Matching a Rule
// from leetcode.com
func CountMatches(items [][]string, ruleKey string, ruleValue string) int {
	var matches = 0

	for _, item := range items {
		if ruleKey == "type" && ruleValue == item[0] {
			matches++
		}

		if ruleKey == "color" && ruleValue == item[1] {
			matches++
		}

		if ruleKey == "name" && ruleValue == item[2] {
			matches++
		}
	}

	return matches
}
