package solutions

import (
	"testing"
)

func TestNumJewelsInStonesCase1(t *testing.T) {
	want := 3

	got := NumJewelsInStones("aA", "aAAbbbb")

	if got != want {
		t.Errorf("NumJewelsInStones() = %d, want %d", got, want)
	}
}
