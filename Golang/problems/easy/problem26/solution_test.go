package solutions

import (
	"testing"
)

func TestRemoveDuplicatesCase1(t *testing.T) {
	want := 5

	arr := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	got := RemoveDuplicates(arr)

	if got != want {
		t.Errorf("RemoveElement() = %d, want %d", got, want)
	}
}
