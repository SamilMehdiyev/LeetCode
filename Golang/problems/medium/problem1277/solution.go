package solutions

// CountSquares function is a solution for the
// Problem #1277 - Count Square Submatrices with All Ones
// from leetcode.com
func CountSquares(matrix [][]int) int {
	n := len(matrix)
	m := len(matrix[0])

	counts := make([][]int, n+1)
	for i := range counts {
		counts[i] = make([]int, m+1)
	}
	total := 0

	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				counts[i+1][j+1] = min(counts[i][j], min(counts[i][j+1], counts[i+1][j])) + 1
			}
			total += counts[i+1][j+1]
		}
	}

	return total
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}
