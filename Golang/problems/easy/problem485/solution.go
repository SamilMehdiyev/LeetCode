package solutions

// FindMaxConsecutiveOnes function is a solution for the
// Problem #485 - Max Consecutive Ones
// from leetcode.com
func FindMaxConsecutiveOnes(nums []int) int {
	maxLength, length := 0, 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			length++
		} else {
			if length > maxLength {
				maxLength = length
			}
			length = 0
		}
	}

	if length > maxLength {
		return length
	}

	return maxLength
}
