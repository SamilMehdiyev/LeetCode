package solutions

// SortedSquare function is a solution for the
// Problem #977 - Squares of a Sorted Array
// from leetcode.com
func SortedSquare(A []int) []int {

	for i := 0; i < len(A); i++ {
		A[i] = A[i] * A[i]
	}

	//selectionSort(&A)
	A = mergeSort(A)

	return A
}

func selectionSort(arr *[]int) {
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

func mergeSort(arr []int) []int {
	if len((arr)) < 2 {
		return arr
	}
	mid := len((arr)) / 2 // divides the current array in 2 parts .

	left := mergeSort(arr[:mid])  // sort the 1st part of array .
	right := mergeSort(arr[mid:]) // sort the 2nd part of array.

	// merge the both parts by comparing elements of both the parts.
	return merge(left, right)
}

func merge(left []int, right []int) []int {
	/*
		stores the starting position of both parts in temporary variables.

		p starting position of the left array
		q starting position of the right array
		capasity is a total length of both arrays (left and right)
	*/
	p, q, capacity := 0, 0, len(left)+len(right)
	array := make([]int, 0, capacity)

	for p < len(left) || q < len(right) {
		if p >= len(left) { //checks if left part comes to an end or not .
			array = append(array, right[q])
			q++
		} else if q >= len(right) { //checks if right part comes to an end or not
			array = append(array, left[p])
			p++
		} else if left[p] < right[q] { //checks which part has smaller element.
			array = append(array, left[p])
			p++
		} else {
			array = append(array, right[q])
			q++
		}
	}

	return array
}
