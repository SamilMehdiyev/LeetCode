package solutions

import (
	"testing"
)

func TestDeleteDuplicatesCase1(t *testing.T) {
	want := []int{1, 2, 3}

	var arr = []int{1, 1, 2, 3, 3}
	node := DeleteDuplicates(arr)

	got := []int{}
	for node != nil {
		got = append(got, node.Val)
		node = node.Next
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}

func TestDeleteDuplicatesCase2(t *testing.T) {
	want := []int{1}

	var arr = []int{1, 1, 1}
	node := DeleteDuplicates(arr)

	got := []int{}
	for node != nil {
		got = append(got, node.Val)
		node = node.Next
	}

	for i := 0; i < len(got); i++ {
		if got[i] != want[i] {
			t.Errorf("Merge() = %d, want %d", got[i], want[i])
		}
	}
}
