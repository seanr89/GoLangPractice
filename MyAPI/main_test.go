/**
* Created using VSCode
* User : Sean
* example file for unit tests!
 */
package main

/*
useful link : https://blog.alexellis.io/golang-writing-unit-tests/
*/

import "testing"

func TestSum(t *testing.T) {
	total := Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
