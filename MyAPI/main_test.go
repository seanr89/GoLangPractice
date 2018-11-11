/**
* Created using VSCode
* User : Sean
* example file for unit tests!
*/
package main

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrec
		t, got: %d, want: %d.", total, 10)
	}
}
