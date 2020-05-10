package solutions

import (
	"testing"
)

func TestCountTripletsCase1(t *testing.T) {
	want := 10
	val := []int{1, 1, 1, 1, 1}

	got := CountTriplets(val)

	if got != want {
		t.Errorf("CountTriplets() = %d, want %d", got, want)
	}
}
