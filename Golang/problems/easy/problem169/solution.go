package solutions

// MajorityElement function is a solution for the
// Problem #169 - Majority Element
// from leetcode.com
func MajorityElement(nums []int) int {
	numsMap := make(map[int]int)
	limit := len(nums) / 2

	for i := 0; i < len(nums); i++ {

		key := nums[i]
		numsMap[key]++

		if numsMap[key] > limit {
			return key
		}
	}
	return 0
}
