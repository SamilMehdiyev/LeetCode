package solutions

import (
	"testing"
)

func TestFindNumbersCase1(t *testing.T) {
	want := 2

	var arr = []int{12, 345, 2, 6, 7896}
	got := FindNumbers(arr)

	if got != want {
		t.Errorf("FindNumbers() = %d, want %d", got, want)
	}
}
