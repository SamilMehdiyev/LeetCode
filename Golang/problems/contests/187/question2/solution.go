package solutions

// KLengthApart function is a solution for the
// Question #2 - Check If All 1's Are at Least Length K Places Away
// from leetcode.com
func KLengthApart(nums []int, k int) bool {

	if nums == nil || len(nums) == 0 {
		return false
	}

	curr, next := 0, 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 && curr == 0 {
			continue
		}

		if curr == 0 {
			curr = i
			next = 0

			for j := i + 1; j < len(nums); j++ {
				if nums[j] == 1 {
					next = j
					break
				}
			}
		}

		if next != curr && next-curr-1 < k {
			return false
		}
	}

	return true
}
