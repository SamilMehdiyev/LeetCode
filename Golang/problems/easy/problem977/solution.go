package solutions

// SortedSquare function is a solution for the
// Problem #977 - Squares of a Sorted Array
// from leetcode.com
func SortedSquare(A []int) []int {

	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	sortArray(&A)

	return A
}

func sortArray(arr *[]int) {
	for i := 0; i < len(*arr); i++ {
		minIndex := i

		for j := i; j < len(*arr); j++ {
			if (*arr)[j] < (*arr)[minIndex] {
				minIndex = j
			}
		}

		temp := (*arr)[i]
		(*arr)[i] = (*arr)[minIndex]
		(*arr)[minIndex] = temp
	}
}
