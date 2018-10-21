/**
* Created using VSCode
* User : Sean
* Basic main class to test hello world tools
*/
package main

import "fmt"
import "./stringUtil"

func main() {

	//initialise a variable to store the output
	var message string
	var a, b, c int
	message = "Hello go world!"
	fmt.Println(message, a, b, c)

	//now lets create a for loop
	for i:=0; i < 5; i++{
		fmt.Println("i is currently : ", i)
	}

	//utilizing package import and method calling
	fmt.Println(stringUtil.Reverse("!oG ,olleH"))
}