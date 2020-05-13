package solutions

// RemoveKdigits2 function is a solution for the
// Problem #402 - Remove K Digits
// from leetcode.com
func RemoveKdigits2(num string, k int) string {

	if k >= len(num) {
		return "0"
	}

	smallest := ""

	for i := 0; i < len(num); i++ {
		for len(smallest) > 0 && smallest[len(smallest)-1] > num[i] && k > 0 {
			smallest = smallest[:len(smallest)-1]
			k--
		}
		smallest += string(num[i])
	}

	for len(smallest) > 1 && int(smallest[0]-'0') == 0 {
		smallest = smallest[1:]
	}

	for i := 0; i < k; i++ {
		smallest = smallest[:len(smallest)-1]
	}

	return smallest
}
