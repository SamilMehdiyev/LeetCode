package solutions

// MaximumWealth function is a solution for the
// Problem #1672. Richest Customer Wealth
// from leetcode.com
func MaximumWealth(accounts [][]int) int {
	wealth := 0
	for i := 0; i < len(accounts); i++ {
		sum := 0
		for j := 0; j < len(accounts[i]); j++ {
			sum += accounts[i][j]
		}
		if sum > wealth {
			wealth = sum
		}
	}
	return wealth
}
