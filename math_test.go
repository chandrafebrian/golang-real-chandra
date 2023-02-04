package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 25 {
		t.Errorf("anjing perhitangan nya salah sat!, harusnya dapet: %d.", total)
	} else {
		println("Testing sum correct bro !")
	}
}
