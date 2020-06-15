package solutions

import (
	"sort"
)

//Problem 787. Cheapest Flights Within K Stops
func findCheapestPrice(n int, flights [][]int, src int, dst int, K int) int {

	cityMap := make(map[int][]City, n)

	for i := 0; i < len(flights); i++ {
		source := flights[i][0]
		destination := flights[i][1]
		price := flights[i][2]

		if cityMap[source] == nil {
			cityMap[source] = []City{}
		}
		cityMap[source] = append(cityMap[source], City{destination, price, 0})
	}

	queue := []City{}
	queue = append(queue, City{src, 0, K + 1})

	for len(queue) > 0 {
		city := queue[0]
		queue = append(queue[1:])

		if city.Destination == dst {
			return city.Price
		}

		if city.Stops > 0 {
			for _, adj := range cityMap[city.Destination] {
				queue = append(queue, City{adj.Destination, city.Price + adj.Price, city.Stops - 1})
			}

			sort.SliceStable(queue, func(i, j int) bool {
				return queue[i].Price < queue[j].Price
			})
		}
	}

	return -1
}

// City struct type
type City struct {
	Destination int
	Price       int
	Stops       int
}
