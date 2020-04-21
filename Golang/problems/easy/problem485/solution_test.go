package solutions

import (
	"testing"
)

func TestFindMaxConsecutiveOnesCase1(t *testing.T) {
	want := 3

	var arr = []int{1, 1, 0, 1, 1, 1}
	got := FindMaxConsecutiveOnes(arr)

	if got != want {
		t.Errorf("FindMaxConsecutiveOnes() = %d, want %d", got, want)
	}
}

func TestFindMaxConsecutiveOnesCase2(t *testing.T) {
	want := 2

	var arr = []int{1, 0, 1, 1, 0, 1}
	got := FindMaxConsecutiveOnes(arr)

	if got != want {
		t.Errorf("FindMaxConsecutiveOnes() = %d, want %d", got, want)
	}
}
