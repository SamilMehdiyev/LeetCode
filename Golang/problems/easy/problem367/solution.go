package solutions

// IsPerfectSquare function is a solution for the
// Problem #367 - Valid Perfect Square
// from leetcode.com
func IsPerfectSquare(num int) bool {

	for i := 1; i*i <= num; i++ {
		if i*i == num {
			return true
		}
	}
	return false
}
