package solutions

// ReplaceElements function is a solution for the
// Problem #1299 - Replace Elements with Greatest Element on Right Side
// from leetcode.com
func ReplaceElements(arr []int) []int {
	max := -1
	for i := len(arr) - 1; 0 <= i; i-- {
		temp := arr[i]
		arr[i] = max

		if temp > max {
			max = temp
		}
	}
	return arr
}
