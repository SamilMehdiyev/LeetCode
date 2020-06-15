package solutions

import (
	"testing"
)

func TestMinSumOfLengthsCase1(t *testing.T) {

	want := 3
	got := MinSumOfLengths([]int{3, 1, 1, 1, 5, 1, 2, 1}, 3)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}
}

func TestMinSumOfLengthsCase2(t *testing.T) {

	want := -1
	got := MinSumOfLengths([]int{5, 5, 4, 4, 5}, 3)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}
}

func TestMinSumOfLengthsCase3(t *testing.T) {

	want := 2
	got := MinSumOfLengths([]int{7, 3, 4, 7}, 7)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}
}

func TestMinSumOfLengthsCase4(t *testing.T) {

	want := -1
	got := MinSumOfLengths([]int{7}, 7)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}
}
func TestMinSumOfLengthsCase5(t *testing.T) {

	want := -1
	got := MinSumOfLengths([]int{3}, 7)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}
}
func TestMinSumOfLengthsCase6(t *testing.T) {

	want := 4
	got := MinSumOfLengths([]int{1, 2, 2, 3, 2, 6, 7, 2, 1, 4, 8}, 5)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}
}
