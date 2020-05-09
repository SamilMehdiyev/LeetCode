package solutions

// CheckStraightLine function is a solution for the
// Problem #1232 - Check If It Is a Straight Line
// from leetcode.com
func CheckStraightLine(coordinates [][]int) bool {

	xDiff := diff(coordinates[0][0], coordinates[1][0])
	yDiff := diff(coordinates[0][1], coordinates[1][1])
	xPrev, yPrev, ratio := coordinates[1][0], coordinates[1][1], divide(xDiff, yDiff)

	for i := 2; i < len(coordinates); i++ {
		if coordinates[i][0] == coordinates[i-1][0] {
			yPrev = coordinates[i][1]
			continue
		}

		if coordinates[i][1] == coordinates[i-1][1] {
			xPrev = coordinates[i][0]
			continue
		}

		xDiff = diff(xPrev, coordinates[i][0])
		yDiff = diff(yPrev, coordinates[i][1])

		if divide(xDiff, yDiff) == ratio {
			ratio = divide(xDiff, yDiff)
			continue
		}

		return false

	}

	return true
}

func diff(xPrev int, x int) int {
	if x-xPrev < 0 {
		return -1 * (x - xPrev)
	}
	return x - xPrev
}

func divide(xDiff int, yDiff int) int {
	if yDiff == 0 {
		return xDiff / 1
	} else if xDiff == 0 {
		return yDiff / 1
	}

	return xDiff / yDiff
}
