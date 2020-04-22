package solutions

import (
	"testing"
)

func TestDuplicateZerosCase1(t *testing.T) {
	want := []int{1, 0, 0, 2, 3, 0, 0, 4}

	var arr = []int{1, 0, 2, 3, 0, 4, 5, 0}
	got := DuplicateZeros(arr)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("DuplicateZeros() = %d, want %d", got[i], want[i])
		}
	}
}
