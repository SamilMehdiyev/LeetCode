package solutions

// SingleNonDuplicate function is a solution for the
// Problem #540 - Remove Duplicates from Sorted Array
// from leetcode.com
func SingleNonDuplicate(nums []int) int {

	if nums == nil || len(nums) == 0 {
		return 0
	} else if len(nums) == 1 {
		return nums[0]
	}

	prev := nums[0]
	for i := 1; i < len(nums)-1; {
		if prev != nums[i] {
			return prev
		}
		prev = nums[i+1]
		i += 2
	}

	return prev
}
