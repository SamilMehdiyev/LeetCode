package solutions

import "strconv"

// SimplifiedFractions function is a solution for the
// Problem #1447 - Simplified Fractions
// from leetcode.com
func SimplifiedFractions(n int) []string {
	i := 1
	j := i + 1
	fractions := []string{}

	for i < n {
		fraction := strconv.Itoa(i) + "/" + strconv.Itoa(j)
		if !isSimplified(i, j) {
			fractions = append(fractions, fraction)
		}
		j++
		if j > n {
			i++
			j = i + 1
		}
	}

	return fractions
}

func isSimplified(i int, j int) bool {

	if i > 1 {
		k := 2

		for ; k <= i; k++ {
			if i%k == 0 && j%k == 0 {
				return true
			}
		}
	}

	return false
}
