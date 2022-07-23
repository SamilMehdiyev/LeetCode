package solutions

import (
	"testing"
)

func TestSubsetXORSum(t *testing.T) {
	got := SubsetXORSum([]int{5, 1, 6})
	want := 28

	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
