package solutions

import (
	"testing"
)

func TestFirstUniqCharCase1(t *testing.T) {
	want := 0

	var arg = "leetcode"
	got := FirstUniqChar(arg)

	if got != want {
		t.Errorf("FirstUniqChar() = %d, want %d", got, want)
	}
}
