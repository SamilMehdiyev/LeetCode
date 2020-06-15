package solutions

import (
	"testing"
)

func TestHasAllCodesCase1(t *testing.T) {
	want := true

	got := HasAllCodes("00000000001011100", 3)

	if got != want {
		t.Errorf("HasAllCodes() = %t, want %t", got, want)
	}
}
