package solutions

import (
	"fmt"
	"math"
	"sort"
)

// KClosest function is a solution for the
// Problem #973 - K Closest Points to Origin
// from leetcode.com
func KClosest(points [][]int, K int) [][]int {
	result := [][]int{}

	pointMap := make(map[int]Point)

	for i := 0; i < len(points); i++ {
		x := float64(points[i][0])
		y := float64(points[i][1])

		distance := math.Sqrt(x*x + y*y)

		pointMap[i] = Point{distance, i, false}
	}

	values := make([]float64, 0, len(pointMap))
	for _, v := range pointMap {
		values = append(values, v.distance)
	}
	sort.Float64s(values)

	for i := 0; i < K; i++ {

		val := values[i]
		for k, v := range pointMap {
			if val == v.distance && !v.isChecked {
				result = append(result, points[k])
				v.isChecked = true
				pointMap[k] = v
				break
			}
		}
		// pointMap = pointMap[:k]
	}

	fmt.Println(points)
	fmt.Println(pointMap)
	fmt.Println(values)
	fmt.Println(result)

	return result
}

type Point struct {
	distance  float64
	idx       int
	isChecked bool
}
