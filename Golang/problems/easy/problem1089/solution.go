package solutions

// DuplicateZeros function is a solution for the
// Problem #1089 - Duplicate Zeros
// from leetcode.com
func DuplicateZeros(arr []int) []int {

	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			for j := len(arr) - 1; i < j; j-- {
				arr[j] = arr[j-1]
			}
			i++
		}
	}

	return arr
}
