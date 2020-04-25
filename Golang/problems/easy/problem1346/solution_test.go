package solutions

import (
	"testing"
)

func TestCheckIfExistCase1(t *testing.T) {
	want := false

	var arr = []int{102, -592, 457, 802, 98, -132, 883, 356, -857, 461,
		-453, 522, 250, 476, 991, 540, -852, -485, -637, 999,
		-803, -691, -880, 881, -584, 750, -124, 745, -909, -892,
		304, -814, 868, 665, 50, -40, 26, -242, -797, -360,
		-918, -741, 88, -933, -93, 360, -738, 833, -191, 563,
		449, 840, 806, -87, -950, 508, 74, -448, -815, -488, 639, -334}
	got := CheckIfExist(arr)

	if got != want {
		t.Errorf("CheckIfExist() = %t, want %t", got, want)
	}
}
