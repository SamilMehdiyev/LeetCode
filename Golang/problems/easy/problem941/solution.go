package solutions

// ValidMountainArray function is a solution for the
// Problem #941 - Valid Mountain Array
// from leetcode.com
func ValidMountainArray(A []int) bool {

	if A != nil && len(A) >= 3 {
		isInc := true
		isDec := true

		for i := 1; i < len(A); i++ {
			if A[i-1] == A[i] {
				return false
			}
			if A[i-1] < A[i] {
				isDec = false
				if isInc == false {
					return false
				}
			} else {
				isInc = false
				if isDec == true {
					return false
				}
			}
		}
		if !isInc && !isDec {
			return true
		}
	}

	return false
}
