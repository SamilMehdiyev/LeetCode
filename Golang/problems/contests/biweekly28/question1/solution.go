package solutions

// FinalPrices function is a solution for the
// Question #1 - Make Two Arrays Equal by Reversing SubArrays
// from leetcode.com
func FinalPrices(prices []int) []int {

	if prices == nil || len(prices) < 2 {
		return prices
	}

	for i := 0; i < len(prices); i++ {
		discount := 0
		for j := i + 1; j < len(prices); j++ {
			if prices[j] <= prices[i] {
				discount = prices[j]
				break
			}
		}
		prices[i] = prices[i] - discount
	}

	return prices
}
