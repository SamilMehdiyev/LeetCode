package solutions

import (
	"math"
	"strconv"
	"strings"
)

// HasAllCodes function is a solution for the
// Question #1 - Check If a String Contains All Binary Codes of Size K
// from leetcode.com
func HasAllCodes(s string, k int) bool {

	binaryCodes := []string{}

	count := int(math.Pow(2, float64(k)))
	for i := 0; i < count; i++ {
		code := strconv.FormatInt(int64(i), 2)

		length := len(code)
		if length < k {
			for i := 0; i < k-length; i++ {
				code = "0" + code
			}
		}
		binaryCodes = append(binaryCodes, code)
	}

	for i := 0; i < len(binaryCodes); i++ {
		if !strings.Contains(s, binaryCodes[i]) {
			return false
		}
	}

	return true
}
