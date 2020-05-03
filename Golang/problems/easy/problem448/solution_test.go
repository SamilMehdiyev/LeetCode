package solutions

import (
	"testing"
)

func TestFindDisappearedNumbersCase1(t *testing.T) {
	want := []int{5, 6}

	var arr = []int{4, 3, 2, 7, 8, 2, 3, 1}
	got := FindDisappearedNumbers(arr)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("FindDisappearedNumbers() = %d, want %d", got[i], want[i])
		}
	}
}
