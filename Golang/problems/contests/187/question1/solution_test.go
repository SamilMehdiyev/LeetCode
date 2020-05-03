package solutions

import (
	"testing"
)

func TestDestCityCase1(t *testing.T) {
	want := "Sao Paulo"

	var arr = [][]string{{"London", "New York"}, {"New York", "Lima"}, {"Lima", "Sao Paulo"}}
	got := DestCity(arr)

	if got != want {
		t.Errorf("DestCity() = %s, want %s", got, want)
	}
}
