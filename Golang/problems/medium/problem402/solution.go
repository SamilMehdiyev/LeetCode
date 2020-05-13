package solutions

// RemoveKdigits function is a solution for the
// Problem #402 - Remove K Digits
// from leetcode.com
func RemoveKdigits(num string, k int) string {
	if k == 0 || len(num) == 0 {
		return num
	}

	smallest := num[1:]
	for i := 1; i < len(num); i++ {
		if num[i-1] == num[i] {
			continue
		}
		smallest = findSmallest(smallest, num[:i]+num[i+1:])
	}

	k--
	num = RemoveKdigits(smallest, k)

	zeroes := 0
	for i := 0; i < len(num); i++ {
		if int(num[i]-'0') != 0 {
			break
		}
		zeroes++
	}

	num = num[zeroes:]

	if len(num) == 0 {
		return "0"
	}

	return num
}

func findSmallest(num1 string, num2 string) string {
	for i := 0; i < len(num1); i++ {
		if int(num1[i]) < int(num2[i]) {
			return num1
		}
	}
	return num2
}
