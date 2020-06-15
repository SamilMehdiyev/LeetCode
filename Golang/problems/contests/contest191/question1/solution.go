package solutions

func maxProduct(nums []int) int {
	maxs := [2]int{0, 0}

	for i := 0; i < len(nums); i++ {
		if maxs[0] <= nums[i] && maxs[1] <= nums[i] {
			maxs[0] = maxs[1]
			maxs[1] = nums[i]
		} else if maxs[0] <= nums[i] && nums[i] <= maxs[1] {
			maxs[0] = nums[i]
		}
	}

	return (maxs[0] - 1) * (maxs[1] - 1)
}
