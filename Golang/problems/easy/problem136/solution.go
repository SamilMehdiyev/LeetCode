package solutions

// SingleNumber function is a solution for the
// Problem #136 - Single Number
// from leetcode.com
func SingleNumber(nums []int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		result = result ^ nums[i]
	}
	return result
}
