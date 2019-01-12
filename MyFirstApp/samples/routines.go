// Package samples contains sample methods
package samples

import "fmt"

func forPrint(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func SampleGoRoutine() {

	forPrint("direct")

	go forPrint("go print")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
}
