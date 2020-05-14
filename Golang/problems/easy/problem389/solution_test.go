package solutions

import (
	"testing"
)

func TestFindTheDifferenceCase1(t *testing.T) {
	want := byte('e')

	got := FindTheDifference("abcd", "abcde")

	if got != want {
		t.Errorf("FirstUniqChar() = %d, want %d", got, want)
	}
}
