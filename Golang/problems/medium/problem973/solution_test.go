package solutions

import (
	"reflect"
	"testing"
)

func TestKClosestCase1(t *testing.T) {
	want := [][]int{{3, 3}, {-2, 4}}

	got := KClosest([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("KClosest() = %d, want %d", got, want)
	}
}

func TestKClosestCase2(t *testing.T) {
	want := [][]int{{0, 1}, {1, 0}}

	got := KClosest([][]int{{0, 1}, {1, 0}}, 2)

	if !reflect.DeepEqual(want, got) {
		t.Errorf("KClosest() = %d, want %d", got, want)
	}
}
