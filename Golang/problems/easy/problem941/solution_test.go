package solutions

import (
	"testing"
)

func TestValidMountainArrayCase1(t *testing.T) {
	want := true

	var arr = []int{0, 3, 2, 1}
	got := ValidMountainArray(arr)

	if got != want {
		t.Errorf("ValidMountainArray() = %t, want %t", got, want)
	}
}

func TestValidMountainArrayCase2(t *testing.T) {
	want := false

	var arr = []int{2, 0, 2}
	got := ValidMountainArray(arr)

	if got != want {
		t.Errorf("ValidMountainArray() = %t, want %t", got, want)
	}
}
