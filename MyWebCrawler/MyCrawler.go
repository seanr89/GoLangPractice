/**
* Created using VSCode
* User : Sean
* Simple web crawler example file to test how to query web pages and the displayed data
 */
package MyCrawler

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	// Make HTTP GET request
    response, err := http.Get("https://www.devdungeon.com/")
    if err != nil {
        log.Fatal(err)
    }
	defer response.Body.Close()
	

}

func crawl(startURL : string) {
}