package solutions

import (
	"testing"
)

func TestRemoveElementCase1(t *testing.T) {
	want := 2

	arr := []int{3, 2, 2, 3}
	val := 3

	got := RemoveElement(arr, val)

	if got != want {
		t.Errorf("RemoveElement() = %d, want %d", got, want)
	}
}
