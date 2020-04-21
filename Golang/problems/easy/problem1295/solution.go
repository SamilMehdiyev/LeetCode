package solutions

// FindNumbers function is a solution for the
// Problem #1295 - Find Numbers with Even Number of Digits
// from leetcode.com
func FindNumbers(nums []int) int {
	eventNums := 0

	for i := 0; i < len(nums); i++ {
		isEven := checkNumberIsEven(nums[i])
		if isEven {
			eventNums++
		}
	}

	return eventNums
}

func checkNumberIsEven(number int) bool {
	digits := 0
	for number > 0 {
		digits++
		number = number / 10
	}

	if digits%2 == 1 {
		return false
	}

	return true
}
