package solutions

import (
	"testing"
)

func TestSortedSquareCase1(t *testing.T) {
	want := []int{0, 1, 9, 16, 100}

	var arr = []int{-4, -1, 0, 3, 10}
	got := SortedSquare(arr)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("SortedSquare() = %d, want %d", got[i], want[i])
		}
	}
}

func TestSortedSquareCase2(t *testing.T) {
	want := []int{4, 9, 9, 49, 121}

	var arr = []int{-7, -3, 2, 3, 11}
	got := SortedSquare(arr)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("SortedSquare() = %d, want %d", got[i], want[i])
		}
	}
}
