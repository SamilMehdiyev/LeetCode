package solutions

import (
	"testing"
)

func TestSubrectangleQueriesCase1(t *testing.T) {

	obj := Constructor([][]int{{1, 1, 1}, {2, 2, 2}, {3, 3, 3}})

	want := 1
	got := obj.GetValue(0, 0)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}

	obj.UpdateSubrectangle(0, 0, 2, 2, 100)

	want = 100
	got = obj.GetValue(0, 0)

	if got != want {
		t.Errorf("GetValue() = %d, want %d", got, want)
	}

	// subrectangleQueries.getValue(0, 0); // return
	// subrectangleQueries.getValue(2, 2); // return 100
	// subrectangleQueries.updateSubrectangle(1, 1, 2, 2, 20);
	// subrectangleQueries.getValue(2, 2); // return 20
}
