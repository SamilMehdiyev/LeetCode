package solutions

// CreateTargetArray function is a solution for the
// Problem #1389. Create Target Array in the Given Order

// from leetcode.com
func CreateTargetArray(nums []int, index []int) []int {
	var target []int

	for i := 0; i < len(index); i++ {
		var idx, value = index[i], nums[i]

		if idx == 0 {
			target = append([]int{value}, target...)
		} else if len(target) == idx {
			target = append(target, value)
		} else if len(target) > 0 && idx < len(target) {
			target = append(target[:idx+1], target[idx:]...)
			target[idx] = value
		}
	}

	return target
}
