package main

import (
	"./crawler"
)

func main() {
	crawler.StartCrawl("https://bbc.co.uk", 3)
	//crawler.StartCrawl("https://www.techcrunch.com/", 3)
}
