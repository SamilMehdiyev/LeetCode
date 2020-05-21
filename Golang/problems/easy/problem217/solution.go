package solutions

// ContainsDuplicate function is a solution for the
// Problem #217 - Contains Duplicate
// from leetcode.com
func ContainsDuplicate(nums []int) bool {
	if nums == nil && len(nums) == 0 {
		return false
	}

	numsMap := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			return true
		}
		numsMap[nums[i]] = true
	}

	return false
}
