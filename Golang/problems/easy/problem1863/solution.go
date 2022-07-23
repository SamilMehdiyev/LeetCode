package solutions

// SubsetXORSum function is a solution for the
// Problem #1863 - Sum of All Subset XOR Totals
// from leetcode.com
func SubsetXORSum(nums []int) int {
	used := make([]bool, len(nums))
	total, length, xor := 0, 0, 0

	backtracking(length, xor, used, nums, &total)

	return total
}

func backtracking(length int, xor int, used []bool, nums []int, total *int) {
	if length == len(nums) {
		return
	}

	for i := length; i < len(nums); i++ {
		if used[i] {
			continue
		}
		used[i] = true
		*total += xor ^ nums[i]
		backtracking(i+1, xor^nums[i], used, nums, total)
		used[i] = false
	}
}
