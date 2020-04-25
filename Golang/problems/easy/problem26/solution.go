package solutions

// RemoveDuplicates function is a solution for the
// Problem #26 - Remove Duplicates from Sorted Array
// from leetcode.com
func RemoveDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	i := 1
	for i < len(nums) {
		if nums[i-1] == nums[i] {
			j := i
			for j < len(nums) && nums[j] == nums[i] {
				j++
			}
			nums = append(nums[:i], nums[j:]...)
			continue
		}
		i++
	}

	return len(nums)
}
