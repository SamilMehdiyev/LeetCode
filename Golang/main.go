package main

import (
	solutions "Golang/problems/easy/problem485"
	"fmt"
)

func main() {
	fmt.Println("Solutions for the Problem from leetcode.com")

	var arr = []int{1, 1, 0, 1, 1, 1}
	fmt.Println(solutions.FindMaxConsecutiveOnes(arr))
}
