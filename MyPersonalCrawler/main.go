package main

import (
	"./crawler"
)

func main() {
	crawler.StartCrawl("https://bbc.co.uk/", 3)
}
