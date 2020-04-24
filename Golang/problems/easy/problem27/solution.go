package solutions

// RemoveElement function is a solution for the
// Problem #27 - Remove Element
// from leetcode.com
func RemoveElement(nums []int, val int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	i := 0
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[:i], nums[i+1:]...)
			continue
		}
		i++
	}

	return len(nums)
}
