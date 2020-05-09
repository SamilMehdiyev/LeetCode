package solutions

import (
	"testing"
)

func TestCheckStraightLineCase1(t *testing.T) {
	want := true

	arr := [][]int{{1, 2}, {2, 3}, {3, 4}, {4, 5}, {5, 6}, {6, 7}}

	got := CheckStraightLine(arr)

	if got != want {
		t.Errorf("CheckStraightLine() = %t, want %t", got, want)
	}
}

func TestCheckStraightLineCase2(t *testing.T) {
	want := true

	arr := [][]int{{-3, -2}, {-1, -2}, {2, -2}, {-2, -2}, {0, -2}}

	got := CheckStraightLine(arr)

	if got != want {
		t.Errorf("CheckStraightLine() = %t, want %t", got, want)
	}
}

func TestCheckStraightLineCase3(t *testing.T) {
	want := true

	arr := [][]int{{0, 1}, {1, 3}, {-4, -7}, {5, 11}}

	got := CheckStraightLine(arr)

	if got != want {
		t.Errorf("CheckStraightLine() = %t, want %t", got, want)
	}
}

func TestCheckStraightLineCase4(t *testing.T) {
	want := true

	arr := [][]int{{1, 1}, {2, 2}, {3, 3}, {4, 4}}

	got := CheckStraightLine(arr)

	if got != want {
		t.Errorf("CheckStraightLine() = %t, want %t", got, want)
	}
}
