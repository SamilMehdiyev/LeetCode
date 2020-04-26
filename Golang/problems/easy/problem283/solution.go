package solutions

// MoveZeroes function is a solution for the
// Problem #283 - Move Zeroes
// from leetcode.com
func MoveZeroes(nums []int) {
	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			temp := nums[nonZeroIndex]
			nums[nonZeroIndex] = nums[i]
			nums[i] = temp

			nonZeroIndex++
		}
	}
}
