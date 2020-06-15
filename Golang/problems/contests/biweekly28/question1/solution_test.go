package solutions

import (
	"reflect"
	"testing"
)

func TestFinalPricesCase1(t *testing.T) {
	want := []int{4, 2, 4, 2, 3}

	got := FinalPrices([]int{8, 4, 6, 2, 3})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("FinalPrices() = %d, want %d", got, want)
	}
}

func TestFinalPricesCase2(t *testing.T) {
	want := []int{}

	got := FinalPrices([]int{})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("FinalPrices() = %d, want %d", got, want)
	}
}

func TestFinalPricesCase3(t *testing.T) {
	want := []int{4}

	got := FinalPrices([]int{4})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("FinalPrices() = %d, want %d", got, want)
	}
}

func TestFinalPricesCase4(t *testing.T) {
	want := []int{1, 2, 3, 4, 5, 6}

	got := FinalPrices([]int{1, 2, 3, 4, 5, 6})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("FinalPrices() = %d, want %d", got, want)
	}
}

func TestFinalPricesCase5(t *testing.T) {
	want := []int{1, 1, 1, 1, 1}

	got := FinalPrices([]int{5, 4, 3, 2, 1})

	if !reflect.DeepEqual(want, got) {
		t.Errorf("FinalPrices() = %d, want %d", got, want)
	}
}
