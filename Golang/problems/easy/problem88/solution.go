package solutions

// Merge function is a solution for the
// Problem #88 - Merge Sorted Array
// from leetcode.com
func Merge(nums1 []int, m int, nums2 []int, n int) []int {

	i, j, nextEmpty := 0, 0, m
	for ; j < n && i < len(nums1); i++ {

		if nums1[i] > nums2[j] {
			for k := len(nums1) - 1; i < k; k-- {
				nums1[k] = nums1[k-1]
			}
			nums1[i] = nums2[j]
			nextEmpty++
			j++
		}
	}

	for ; j < n; j++ {
		nums1[nextEmpty] = nums2[j]
		nextEmpty++
	}

	return nums1
}
