package solutions

import (
	"testing"
)

func TestIsPerfectSquareCase1(t *testing.T) {
	want := false
	val := 3

	got := IsPerfectSquare(val)

	if got != want {
		t.Errorf("IsPerfectSquare() = %t, want %t", got, want)
	}
}

func TestIsPerfectSquareCase2(t *testing.T) {
	want := false
	val := 14

	got := IsPerfectSquare(val)

	if got != want {
		t.Errorf("IsPerfectSquare() = %t, want %t", got, want)
	}
}

func TestIsPerfectSquareCase3(t *testing.T) {
	want := true
	val := 16

	got := IsPerfectSquare(val)

	if got != want {
		t.Errorf("IsPerfectSquare() = %t, want %t", got, want)
	}
}
