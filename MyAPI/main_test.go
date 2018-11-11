package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorreclst, got: %d, want: %d.", total, 10)
	}
}
