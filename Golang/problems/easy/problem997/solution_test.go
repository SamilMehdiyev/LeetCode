package solutions

import (
	"testing"
)

func TestFindJudgeCase1(t *testing.T) {
	want := 3

	n := 4
	trust := [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}
	got := FindJudge(n, trust)

	if got != want {
		t.Errorf("FindJudge() = %d, want %d", got, want)
	}
}

func TestFindJudgeCase2(t *testing.T) {
	want := -1

	n := 3
	trust := [][]int{{1, 2}, {2, 3}}
	got := FindJudge(n, trust)

	if got != want {
		t.Errorf("FindJudge() = %d, want %d", got, want)
	}
}

func TestFindJudgeCase3(t *testing.T) {
	want := -1

	n := 3
	trust := [][]int{{1, 3}, {2, 3}, {3, 1}}
	got := FindJudge(n, trust)

	if got != want {
		t.Errorf("FindJudge() = %d, want %d", got, want)
	}
}
