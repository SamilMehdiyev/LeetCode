package solutions

// FindDisappearedNumbers function is a solution for the
// Problem #448 - Find All Numbers Disappeared in an Array
// from leetcode.com
func FindDisappearedNumbers(nums []int) []int {

	for i := 0; i < len(nums); i++ {
		index := abs(nums[i]) - 1
		if nums[index] > 0 {
			nums[index] = -1 * nums[index]
		}
	}

	arr := []int{}

	for i := 0; i < len(nums); i++ {
		if nums[i] > 0 {
			arr = append(arr, i+1)
		}
	}

	return arr
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}

	return x
}
