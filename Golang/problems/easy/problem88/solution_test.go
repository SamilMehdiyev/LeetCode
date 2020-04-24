package solutions

import (
	"testing"
)

func TestMergeCase1(t *testing.T) {
	want := []int{1, 2, 2, 3, 5, 6}

	var arr1 = []int{1, 2, 3, 0, 0, 0}
	var m = 3
	var arr2 = []int{2, 5, 6}
	var n = 3
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase2(t *testing.T) {
	want := []int{1}

	var arr1 = []int{1}
	var m = 1
	var arr2 = []int{}
	var n = 0
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase3(t *testing.T) {
	want := []int{1}

	var arr1 = []int{0}
	var m = 0
	var arr2 = []int{1}
	var n = 1
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase4(t *testing.T) {
	want := []int{1, 2}

	var arr1 = []int{1, 0}
	var m = 1
	var arr2 = []int{2}
	var n = 1
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase5(t *testing.T) {
	want := []int{1, 2}

	var arr1 = []int{2, 0}
	var m = 1
	var arr2 = []int{1}
	var n = 1
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase6(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6}

	var arr1 = []int{4, 0, 0, 0, 0, 0}
	var m = 1
	var arr2 = []int{1, 2, 3, 5, 6}
	var n = 5
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase7(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6}

	var arr1 = []int{4, 5, 6, 0, 0, 0}
	var m = 3
	var arr2 = []int{1, 2, 3}
	var n = 3
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase8(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6}

	var arr1 = []int{1, 2, 4, 5, 6, 0}
	var m = 5
	var arr2 = []int{3}
	var n = 1
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestMergeCase9(t *testing.T) {
	want := []int{-1, 0, 0, 1, 2, 2, 3, 3, 3}

	var arr1 = []int{-1, 0, 0, 3, 3, 3, 0, 0, 0}
	var m = 6
	var arr2 = []int{1, 2, 2}
	var n = 3
	got := Merge(arr1, m, arr2, n)

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}
