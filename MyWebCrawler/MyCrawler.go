/**
* Created using VSCode
* User : Sean
* Simple web crawler example file to test how to query web pages and the displayed data
 */
package main

/*
useful urls
https://www.devdungeon.com/content/web-scraping-go
https://tour.golang.org/concurrency/10
http://edmundmartin.com/writing-a-web-crawler-in-golang/
https://jdanger.com/build-a-web-crawler-in-go.html
*/

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func main() {

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Make request using timeout client!
	response, err := client.Get("https://www.devdungeon.com/")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	// Create output file if you want to store the site in a file
	// outFile, err := os.Create("output.html")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer outFile.Close()

	// Get the response body as a string
	dataInBytes, err := ioutil.ReadAll(response.Body)
	pageContent := string(dataInBytes)

	// Find a substr
	titleStartIndex := strings.Index(pageContent, "<title>")
	if titleStartIndex == -1 {
		fmt.Println("No title element found")
		os.Exit(0)
	}

	// Copy data from the response to standard output
	n, err := io.Copy(os.Stdout, response.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Number of bytes copied to STDOUT:", n)
}

func crawl(startURL string) {
}
