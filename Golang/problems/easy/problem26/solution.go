package solutions

// RemoveDuplicates function is a solution for the
// Problem #26 - Remove Duplicates from Sorted Array
// from leetcode.com
func RemoveDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	length, i := 1, 1
	for i < len(nums) {
		if nums[i-1] != nums[i] {
			nums[length] = nums[i]
			length++
		}
		i++
	}

	return length
}
