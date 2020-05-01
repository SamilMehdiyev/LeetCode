package solutions

// FirstBadVersion function is a solution for the
// Problem #278 - First Bad Version
// from leetcode.com
func FirstBadVersion(n int) int {
	start, end := 1, n
	firstBadVersion := 0

	for start < end {
		mid := (start + end) / 2
		result := isBadVersion(mid)

		if result {
			end = mid
			firstBadVersion = mid
		} else {
			start = mid + 1
		}
	}

	return firstBadVersion
}

func isBadVersion(version int) bool {
	badVersion := 4
	if version == badVersion {
		return true
	}

	return false
}
