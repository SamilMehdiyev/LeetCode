package solutions

import (
	"sort"
)

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	maxH, maxW := 0, 0

	maxH = findMaxDistance(horizontalCuts, h)
	maxW = findMaxDistance(verticalCuts, w)

	return ((maxH % 1_000_000_007) * (maxW % 1_000_000_007)) % 1_000_000_007
}

func findMaxDistance(cuts []int, value int) int {
	sort.Ints(cuts)

	prev, distance, max := 0, 0, 0
	for i := 0; i < len(cuts); i++ {
		distance = cuts[i] - prev

		if distance > max {
			max = distance
		}

		prev = cuts[i]
	}

	if value-prev > max {
		max = value - prev
	}

	return max
}
