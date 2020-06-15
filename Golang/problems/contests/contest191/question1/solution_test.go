package solutions

import (
	"testing"
)

func TestMaxProductCase1(t *testing.T) {
	want := 12

	got := maxProduct([]int{3, 4, 5, 2})

	if got != want {
		t.Errorf("HasAllCodes() = %d, want %d", got, want)
	}
}

func TestMaxProductCase2(t *testing.T) {
	want := 16

	got := maxProduct([]int{1, 5, 4, 5})

	if got != want {
		t.Errorf("HasAllCodes() = %d, want %d", got, want)
	}
}

func TestMaxProductCase3(t *testing.T) {
	want := 12

	got := maxProduct([]int{3, 7})

	if got != want {
		t.Errorf("HasAllCodes() = %d, want %d", got, want)
	}
}
