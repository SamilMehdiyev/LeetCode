package solutions

// CanBeEqual function is a solution for the
// Question #1 - Make Two Arrays Equal by Reversing SubArrays
// from leetcode.com
func CanBeEqual(target []int, arr []int) bool {

	diff := 0
	numMap := make(map[int]int)

	for i := 0; i < len(target); i++ {
		numMap[target[i]]++
		diff++
	}

	for i := 0; i < len(arr); i++ {
		if numMap[arr[i]] > 0 {
			numMap[arr[i]]--
			diff--
		}
	}

	if diff > 0 {
		return false
	}

	return true
}
