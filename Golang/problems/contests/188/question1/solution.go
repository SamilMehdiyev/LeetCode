package solutions

// BuildArray function is a solution for the
// Question #1 - Build an Array With Stack Operations
// from leetcode.com
func BuildArray(target []int, n int) []string {
	operatios := []string{}

	index := 0
	for i := 1; i <= n; i++ {
		if index >= len(target) {
			break
		} else if i < target[index] {
			operatios = append(operatios, "Push", "Pop")
		} else if i == target[index] {
			operatios = append(operatios, "Push")
			index++
		}
	}
	return operatios
}
