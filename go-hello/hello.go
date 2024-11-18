package main

import "fmt"
// import random quotes from rsc.io/quote
import "rsc.io/quote"

func main() {
    fmt.Println("Hello, World!")
	fmt.Println(quote.Go())
}