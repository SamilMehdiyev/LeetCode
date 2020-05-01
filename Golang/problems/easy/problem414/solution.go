package solutions

// ThirdMax function is a solution for the
// Problem #414 - Third Maximum Number
// from leetcode.com
func ThirdMax(nums []int) int {

	const UintSize = 32 << (^uint(0) >> 32 & 1) // 32 or 64

	const (
		MaxInt = 1<<(UintSize-1) - 1 // 1<<31 - 1 or 1<<63 - 1
		MinInt = -MaxInt - 1         // -1 << 31 or -1 << 63
	)

	// Max 3 numbers a > b > c
	a, b, c := MinInt, MinInt, MinInt
	prev := -1

	for i := 0; i < len(nums); i++ {
		if prev == nums[i] {
			continue
		}

		if nums[i] > a {
			c = b
			b = a
			a = nums[i]

			continue
		}

		if a > nums[i] && nums[i] > b {
			c = b
			b = nums[i]

			continue
		}

		if b > nums[i] && nums[i] > c {
			c = nums[i]

			continue
		}

		prev = nums[i]
	}

	if c != MinInt {
		return c
	}

	return a
}
