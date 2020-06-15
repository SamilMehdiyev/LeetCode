package solutions

func checkIfPrerequisite(n int, prerequisites [][]int, queries [][]int) []bool {

	adjList := make([][]int, n)
	seen := make([]int, n)

	for i := 0; i < len(adjList); i++ {
		adjList[i] = []int{}
	}

	for _, v := range prerequisites {
		x, y := v[0], v[1]
		adjList[x] = append(adjList[x], y)
	}

	results := []bool{}

	for _, v := range queries {
		x, y := v[0], v[1]

		if !checkCourse(x, y, &adjList, &seen) {
			results = append(results, false)
		} else {
			results = append(results, true)
		}
	}

	return results
}

func checkCourse(x int, y int, adjList *[][]int, seen *[]int) bool {
	(*seen)[x] = 1

	if x == y {
		return true
	}

	for _, v := range (*adjList)[x] {
		if (*seen)[v] == 0 && checkCourse(v, y, adjList, seen) {
			return true
		}
	}

	return false
}
