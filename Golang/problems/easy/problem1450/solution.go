package solutions

// BusyStudent function is a solution for the
// Problem #1450 - Number of Students Doing Homework at a Given Time
// from leetcode.com
func BusyStudent(startTime []int, endTime []int, queryTime int) int {
	if startTime == nil || len(startTime) == 0 {
		return 0
	} else if endTime == nil || len(endTime) == 0 {
		return 0
	}

	students := 0
	length := 0

	if len(startTime) < len(endTime) {
		length = len(startTime)
	} else {
		length = len(endTime)
	}

	for i := 0; i < length; i++ {
		if startTime[i] <= queryTime && queryTime <= endTime[i] {
			students++
		}
	}

	return students
}
