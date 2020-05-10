package solutions

import (
	"testing"
)

func TestBuildArrayCase1(t *testing.T) {
	want := []string{"Push", "Push", "Pop", "Push"}

	var target = []int{1, 3}
	var n = 3
	got := BuildArray(target, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("BuildArray() = %s, want %s", got[i], want[i])
		}
	}
}

func TestBuildArrayCase2(t *testing.T) {
	want := []string{"Push", "Pop", "Push", "Push", "Push"}

	var target = []int{2, 3, 4}
	var n = 4
	got := BuildArray(target, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("BuildArray() = %s, want %s", got[i], want[i])
		}
	}
}

func TestBuildArrayCase3(t *testing.T) {
	want := []string{"Push", "Push"}

	var target = []int{1, 2}
	var n = 4
	got := BuildArray(target, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("BuildArray() = %s, want %s", got[i], want[i])
		}
	}
}
