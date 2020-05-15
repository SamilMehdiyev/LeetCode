package solutions

// MaxSubarraySumCircular function is a solution for the
// Problem #918 - Maximum Sum Circular Subarray
// from leetcode.com
func MaxSubarraySumCircular(A []int) int {
	sumMax := -30000
	sumMin := 30000

	currMax, currMin, sum := 0, 0, 0

	for i := 0; i < len(A); i++ {
		currMax = max(currMax+A[i], A[i])
		sumMax = max(currMax, sumMax)

		currMin = min(currMin+A[i], A[i])
		sumMin = min(currMin, sumMin)

		sum += A[i]
	}

	if sum != sumMin {
		return max(sum-sumMin, sumMax)
	}

	return sumMax
}

func max(num1 int, num2 int) int {
	if num1 < num2 {
		return num2
	}
	return num1
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
