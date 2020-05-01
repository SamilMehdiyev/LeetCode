package solutions

import (
	"testing"
)

func TestThirdMaxCase1(t *testing.T) {
	want := 1

	var arr = []int{2, 2, 3, 1}
	got := ThirdMax(arr)

	if got != want {
		t.Errorf("ThirdMax() = %d, want %d", got, want)
	}
}

func TestThirdMaxCase2(t *testing.T) {
	want := 2

	var arr = []int{1, 2}
	got := ThirdMax(arr)

	if got != want {
		t.Errorf("ThirdMax() = %d, want %d", got, want)
	}
}

func TestThirdMaxCase3(t *testing.T) {
	want := 5

	var arr = []int{5, 2, 2}
	got := ThirdMax(arr)

	if got != want {
		t.Errorf("ThirdMax() = %d, want %d", got, want)
	}
}

func TestThirdMaxCase4(t *testing.T) {
	want := 4

	var arr = []int{2, 2, 3, 1, 4, 5, 6, 3, 4, 2}
	got := ThirdMax(arr)

	if got != want {
		t.Errorf("ThirdMax() = %d, want %d", got, want)
	}
}

func TestThirdMaxCase5(t *testing.T) {
	want := 2

	var arr = []int{1, 2, 2, 5, 3, 5}
	got := ThirdMax(arr)

	if got != want {
		t.Errorf("ThirdMax() = %d, want %d", got, want)
	}
}
