package solutions

import (
	"testing"
)

func TestFloodFillCase1(t *testing.T) {
	image := [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}
	want := [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}

	got := FloodFill(image, 1, 1, 2)

	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("FloodFill() = %d, want %d", got[i], want[i])
			}
		}
	}
}
