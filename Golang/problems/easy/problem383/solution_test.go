package solutions

import (
	"testing"
)

func TestCanConstructCase1(t *testing.T) {
	want := true

	note, magazine := "aa", "aab"
	got := CanConstruct(note, magazine)

	if got != want {
		t.Errorf("CanConstruct() = %t, want %t", got, want)
	}
}

func TestCanConstructCase2(t *testing.T) {
	want := false

	note, magazine := "aa", "ab"
	got := CanConstruct(note, magazine)

	if got != want {
		t.Errorf("CanConstruct() = %t, want %t", got, want)
	}
}
