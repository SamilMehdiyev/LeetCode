package solutions

import (
	"testing"
)

func TestMajorityElementCase1(t *testing.T) {
	want := 2

	got := MajorityElement([]int{2, 2, 1, 1, 1, 2, 2})

	if got != want {
		t.Errorf("MajorityElement() = %d, want %d", got, want)
	}
}
