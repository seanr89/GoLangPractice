// /**
// * Created using VSCode
// * User : Sean
// * Simple web crawler example file to test how to query web pages and the displayed data
//  */
// package main

// /*
// useful urls
// https://www.devdungeon.com/content/web-scraping-go
// https://tour.golang.org/concurrency/10
// http://edmundmartin.com/writing-a-web-crawler-in-golang/
// https://jdanger.com/build-a-web-crawler-in-go.html
// */

// import (
// 	"fmt"
// 	"io"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"os"
// 	"strings"
// 	"time"
// )

// func main() {

// 	// Create HTTP client with timeout
// 	client := &http.Client{
// 		Timeout: 30 * time.Second,
// 	}

// 	// Make request using timeout client!
// 	response, err := client.Get("https://www.devdungeon.com/")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer response.Body.Close()

// 	// Create output file if you want to store the site in a file
// 	// outFile, err := os.Create("output.html")
// 	// if err != nil {
// 	// 	log.Fatal(err)
// 	// }
// 	// defer outFile.Close()

// 	// Get the response body as a string
// 	dataInBytes, err := ioutil.ReadAll(response.Body)
// 	pageContent := string(dataInBytes)

// 	// Find a substr
// 	titleStartIndex := strings.Index(pageContent, "Clearance")
// 	if titleStartIndex == -1 {
// 		fmt.Println("No matching content found")
// 		os.Exit(0)
// 	}

// 	// Copy data from the response to standard output
// 	n, err := io.Copy(os.Stdout, response.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	log.Println("Number of bytes copied to STDOUT:", n)
// }

// func retrieve(uri string) { // This func(tion) takes a parameter and the
// 	// format for a function parameter definition is
// 	// to say what the name of the parameter is and then
// 	// the type.
// 	// So here we're expecting to be given a
// 	// string that we'll refer to as 'uri'
// 	resp, err := http.Get(uri)
// 	if err != nil { // This is the way error handling typically works in Go.
// 		return // It's a bit verbose but it works.
// 	}
// 	defer resp.Body.Close() // Important: we need to close the resource we opened
// 	// (the TCP connection to some web server and our reference
// 	// to the stream of data it sends us).
// 	// `defer` delays an operation until the function ends.
// 	// It's basically the same as if you'd moved the code
// 	// you're deferring to the very last line of the func.

// 	body, _ := ioutil.ReadAll(resp.Body) // I'm assigning the err to _ 'cause
// 	// I don't care about it but Go will whine
// 	fmt.Println(string(body)) // if I name it and don't use it
// }
