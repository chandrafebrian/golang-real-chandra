package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 25 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 121)
	} else {
		println("Testing sum correct bro !")
	}
}
