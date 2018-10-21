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

	//testing the executing of method calls from the file itself as well
	doDBOperations()
}

func connectToDB () {
    fmt.Println( "ok, connected to db" )
}

func disconnectFromDB () {
    fmt.Println( "ok, disconnected from db" )
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