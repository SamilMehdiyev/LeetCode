package solutions

import (
	"testing"
)

func TestFirstBadVersionCase1(t *testing.T) {
	want := 4

	got := FirstBadVersion(5)

	if got != want {
		t.Errorf("ThirdMax() = %d, want %d", got, want)
	}
}
