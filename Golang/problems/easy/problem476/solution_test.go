package solutions

import (
	"testing"
)

func TestFindComplementCase1(t *testing.T) {
	want := 2

	var number = 5
	got := FindComplement(number)

	if got != want {
		t.Errorf("FindComplement() = %d, want %d", got, want)
	}
}

func TestFindComplementCase2(t *testing.T) {
	want := 1

	var number = 2
	got := FindComplement(number)

	if got != want {
		t.Errorf("FindComplement() = %d, want %d", got, want)
	}
}

func TestFindComplementCase3(t *testing.T) {
	want := 0

	var number = 1
	got := FindComplement(number)

	if got != want {
		t.Errorf("FindComplement() = %d, want %d", got, want)
	}
}
