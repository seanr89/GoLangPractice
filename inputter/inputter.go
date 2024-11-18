package inputter

import "fmt"

// Hello returns a greeting for the named person.
func AskQuestion() {
    // Println function is used to 
    // display output in the next line 
    fmt.Println("Enter Your First Name: ") 
  
    // var then variable name then variable type 
    var first string 
  
    // Taking input from user 
    fmt.Scanln(&first) 
    fmt.Println("Enter Second Last Name: ") 
    var second string 
    fmt.Scanln(&second) 
  
    // Print function is used to 
    // display output in the same line 
    fmt.Print("Your Full Name is: ") 
  
    // Addition of two string 
    fmt.Println(first + " " + second) 
}