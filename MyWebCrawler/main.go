package main

import (
	"./crawler"
	"github.com/PuerkitoBio/goquery"
)

/**
 */
type DummyParser struct {
}

func (d DummyParser) ParsePage(doc *goquery.Document) crawler.ScrapeResult {
	data := crawler.ScrapeResult{}
	data.Title = doc.Find("title").First().Text()
	data.H1 = doc.Find("h1").First().Text()
	return crawler.ScrapeResult{}
}

func main() {
	//d := DummyParser{}
	//crawler.Crawl("https://bbc.co.uk/", d, 3)
	//crawler.Crawl("https://vsorelease.azurewebsites.net/", d, 4)

	crawler.StartCrawl("https://bbc.co.uk/", 3)
}
