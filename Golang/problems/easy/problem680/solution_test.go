package solutions

import (
	"testing"
)

func TestValidPalindromeCase1(t *testing.T) {
	want := true

	var str = "abca"
	got := ValidPalindrome(str)

	if got != want {
		t.Errorf("ValidPalindrome() = %t, want %t", got, want)
	}
}

func TestValidPalindromeCase2(t *testing.T) {
	want := true

	var str = "aba"
	got := ValidPalindrome(str)

	if got != want {
		t.Errorf("ValidPalindrome() = %t, want %t", got, want)
	}
}

func TestValidPalindromeCase3(t *testing.T) {
	want := true

	var str = "ebcbbececabbacecbbcbe"
	got := ValidPalindrome(str)

	if got != want {
		t.Errorf("ValidPalindrome() = %t, want %t", got, want)
	}
}

func TestValidPalindromeCase4(t *testing.T) {
	want := true

	var str = "cbbcc"
	got := ValidPalindrome(str)

	if got != want {
		t.Errorf("ValidPalindrome() = %t, want %t", got, want)
	}
}
