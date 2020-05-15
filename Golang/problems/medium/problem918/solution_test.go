package solutions

import (
	"testing"
)

func TestMaxSubarraySumCircularCase1(t *testing.T) {
	want := 3

	arr := []int{1, -2, 3, -2}

	got := MaxSubarraySumCircular(arr)

	if got != want {
		t.Errorf("MaxSubarraySumCircular() = %d, want %d", got, want)
	}
}

func TestMaxSubarraySumCircularCase2(t *testing.T) {
	want := 10

	arr := []int{5, -3, 5}

	got := MaxSubarraySumCircular(arr)

	if got != want {
		t.Errorf("MaxSubarraySumCircular() = %d, want %d", got, want)
	}
}
