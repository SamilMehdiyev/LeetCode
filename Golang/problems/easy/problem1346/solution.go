package solutions

// CheckIfExist function is a solution for the
// Problem #1346 - Check if N and its Double Exist
// from leetcode.com
func CheckIfExist(arr []int) bool {

	for i := 0; i < len(arr); i++ {
		if arr[i]%2 == 1 || arr[i]%2 == -1 {
			continue
		}
		val := arr[i] / 2
		for j := 0; j < len(arr); j++ {
			if i != j && arr[j] == val {
				return true
			}
		}
	}
	return false
}
