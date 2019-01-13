/**
* Created using VSCode
* User : Sean
* Basic main class to test hello world tools
 */
package main

import "fmt"
import "./stringUtil"

import "./samples"

func main() {

	//initialise a variable to store the output
	var message string
	message = "Hello go world!"
	fmt.Println(message)

	//now lets create a for loop
	// for i := 0; i < 5; i++ {
	// 	fmt.Println("i is currently : ", i)
	// }

	//utilizing package import and method calling
	fmt.Println(stringUtil.Reverse("!oG ,olleH"))

	//testing the executing of method calls from the file itself as well
	//doDBOperations()

	//example of how constants work
	//constSamp.ShowConstants()

	//variadicSum(1, 2, 3, 4)

	//Example of a go thread execution
	samples.SampleGoRoutine()

	//Execute a set of go worker threads and pass jobs to each
	samples.RunWorkers(2, 15)
}

func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

/*
testing with deferred operations calls (to be executed just before returning from the current function block)
*/
func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB() //function called here with defer
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return //terminate the program

	// deferred function executed here just before actually returning, even if there is a return or abnormal termination before
}

// Variadic functions can be called with any number of trailing arguments.
func variadicSum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}
