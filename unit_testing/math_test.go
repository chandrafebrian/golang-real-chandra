package unit_testing

import "testing"

func Sum(x int, y int) int {
	return x * y
}

func TestSum(t *testing.T) {
	total := Sum(5, 5) //total adalah var default menampung fungsi Sum(parameter int, int)
	if total != 35 {
		t.Errorf(" perhitangan nya salah sat!, harusnya hasil: %d.", total)
	} else {
		println("Testing sum correct bro !")
	}
}
