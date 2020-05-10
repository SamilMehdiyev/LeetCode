package solutions

// CountTriplets function is a solution for the
// Question #2 - Count Triplets That Can Form Two Arrays of Equal XOR
// from leetcode.com
func CountTriplets(arr []int) int {
	if len(arr) < 2 {
		return 0
	}

	i, j, k := 0, 1, 1
	a, b := arr[i], arr[j]
	count := 0

	for ; i < len(arr)-2; k++ {
		if k >= len(arr) {
			j++
			if j >= len(arr) {
				i++
				j = i + 1
			}
			k = j
		}

		a = calculate(arr, i, j-1)
		b = calculate(arr, j, k)

		if a == b {
			count++
		}
	}

	return count
}

func calculate(arr []int, from int, to int) int {
	b := arr[from]
	for i := from + 1; i <= to; i++ {
		b ^= arr[i]
	}
	return b
}
