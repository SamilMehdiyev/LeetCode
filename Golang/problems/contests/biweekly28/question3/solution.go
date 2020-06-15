package solutions

import (
	"sort"
)

// MinSumOfLengths function is a solution for the
// Question #3 - Make Two Arrays Equal by Reversing SubArrays
// from leetcode.com
func MinSumOfLengths(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return 0
	}

	result := -1
	sum, length, lengths := 0, 0, []int{}
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		length++
		if sum > target {
			i = i - length + 1
			sum = 0
			length = 0
		} else if sum == target {
			if length > 0 {
				lengths = append(lengths, length)
			}
			sum = 0
			length = 0
		}
	}
	sort.Ints(lengths)

	if len(lengths) >= 2 {
		result = lengths[0] + lengths[1]
	}

	return result
}

func minSumOfLengths(arr []int, target int) int {
	if arr == nil || len(arr) == 0 {
		return -1
	}

	const MaxValue int = 1000000
	arrLength := len(arr)
	lengths := make([]int, arrLength)
	for i := 0; i < len(lengths); i++ {
		lengths[i] = MaxValue
	}

	sum, start, result, minLength, length := 0, 0, MaxValue, MaxValue, 0

	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		length++
		if sum > target {
			for sum > target {
				sum -= arr[start]
				length--
				start++
			}
		}

		if sum == target {
			if start > 0 && lengths[start-1] != MaxValue {
				result = min(result, lengths[start-1]+length)
			}
			minLength = min(minLength, length)
		}
		lengths[i] = minLength
	}

	if result != MaxValue {
		return result
	}
	return -1
}

func min(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}
	return num2
}
