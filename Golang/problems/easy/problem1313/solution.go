package solutions

// DecompressRLElist function is a solution for the
// Problem #1313. Decompress Run-Length Encoded List
// from leetcode.com
func DecompressRLElist(nums []int) []int {
	var result []int

	for i := 0; i < len(nums); i += 2 {
		var freq, val = nums[i], nums[i+1]
		for ; freq > 0; freq-- {
			result = append(result, val)
		}
	}
	return result
}
