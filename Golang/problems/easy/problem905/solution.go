package solutions

// SortArrayByParity function is a solution for the
// Problem #905 - Sort Array By Parity
// from leetcode.com
func SortArrayByParity(A []int) []int {
	nextEvenIndex := 0

	for i := len(A) - 1; 0 < i && nextEvenIndex < i; {
		if A[i]%2 == 0 {
			if A[nextEvenIndex]%2 == 1 {
				A[nextEvenIndex], A[i] = A[i], A[nextEvenIndex]
				i--
			}
			nextEvenIndex++
		} else {
			i--
		}
	}

	return A
}
