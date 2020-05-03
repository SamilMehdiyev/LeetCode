package solutions

import (
	"testing"
)

func TestKLengthApartCase1(t *testing.T) {
	want := true

	var nums, k = []int{1, 0, 0, 0, 1, 0, 0, 1}, 2
	got := KLengthApart(nums, k)

	if got != want {
		t.Errorf("KLengthApart() = %t, want %t", got, want)
	}
}

func TestKLengthApartCase2(t *testing.T) {
	want := false

	var nums, k = []int{1, 0, 0, 1, 0, 1}, 2
	got := KLengthApart(nums, k)

	if got != want {
		t.Errorf("KLengthApart() = %t, want %t", got, want)
	}
}

func TestKLengthApartCase3(t *testing.T) {
	want := true

	var nums, k = []int{1, 1, 1, 1, 1}, 0
	got := KLengthApart(nums, k)

	if got != want {
		t.Errorf("KLengthApart() = %t, want %t", got, want)
	}
}

func TestKLengthApartCase4(t *testing.T) {
	want := true

	var nums, k = []int{0, 1, 0, 1}, 1
	got := KLengthApart(nums, k)

	if got != want {
		t.Errorf("KLengthApart() = %t, want %t", got, want)
	}
}

func TestKLengthApartCase5(t *testing.T) {
	want := true

	var nums, k = []int{1, 0, 0, 0, 0}, 3
	got := KLengthApart(nums, k)

	if got != want {
		t.Errorf("KLengthApart() = %t, want %t", got, want)
	}
}
