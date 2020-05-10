package solutions

// FindJudge function is a solution for the
// Problem #997 - Find the Town Judge
// from leetcode.com
func FindJudge(N int, trust [][]int) int {

	if len(trust) == 0 {
		return N
	}

	judges := make([]int, N+1)

	for i := 0; i < len(trust); i++ {
		personA := trust[i][0]
		personB := trust[i][1]

		judges[personA]--
		judges[personB]++
	}

	for i := 0; i < len(judges); i++ {
		if judges[i] == N-1 {
			return i
		}
	}

	return -1
}
