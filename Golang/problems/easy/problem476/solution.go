package solutions

// FindComplement function is a solution for the
// Problem #476 - Number Complement
// from leetcode.com
func FindComplement(num int) int {

	bits := []byte{}
	for num > 0 {
		reminder := num % 2
		bits = append(bits, byte(reminder)^1)

		num = num / 2
	}

	value, power := 0, 1
	for i := 0; i < len(bits); i++ {
		value += int(bits[i]) * power
		power *= 2
	}

	return value
}
