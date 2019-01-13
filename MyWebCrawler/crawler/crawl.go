package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type ScrapeResult struct {
	URL   string
	Title string
	H1    string
}

/**
Parser interfaces and initialises call for  goQuery to read the provided document
*/
type Parser interface {
	ParsePage(*goquery.Document) ScrapeResult
}

func getRequest(url string) (*http.Response, error) {
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

/*
Extract links for the url goquery document
return as a list of strings
*/
func extractLinks(doc *goquery.Document) []string {
	foundUrls := []string{}
	if doc != nil {
		doc.Find("a").Each(func(i int, s *goquery.Selection) {
			res, _ := s.Attr("href")
			foundUrls = append(foundUrls, res)
		})
		return foundUrls
	}
	return foundUrls
}

func resolveRelative(baseURL string, hrefs []string) []string {
	internalUrls := []string{}

	for _, href := range hrefs {
		if strings.HasPrefix(href, baseURL) {
			internalUrls = append(internalUrls, href)
		}

		if strings.HasPrefix(href, "/") {
			resolvedURL := fmt.Sprintf("%s%s", baseURL, href)
			internalUrls = append(internalUrls, resolvedURL)
		}
	}
	return internalUrls
}

/*
Executing the processing of inidividual page contents
*/
func crawlPage(baseURL, targetURL string, parser Parser, token chan struct{}) ([]string, ScrapeResult) {

	token <- struct{}{}
	fmt.Println("Requesting: ", targetURL)
	resp, _ := getRequest(targetURL)
	<-token

	doc, _ := goquery.NewDocumentFromResponse(resp)
	pageResults := parser.ParsePage(doc)
	links := extractLinks(doc)
	foundUrls := resolveRelative(baseURL, links)

	return foundUrls, pageResults
}

/*
parseStartURL output the scheme and the host of the site
returns a output string to console with the scheme and host
*/
func parseStartURL(u string) string {
	// use url package to parse the url from the string
	parsed, _ := url.Parse(u)
	return fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)
}

/**
readAndPrintAllUrls task is to read in each url and output the string content
*/
func readAddPrintAllUrls(data []string) {
	for i := range data {
		fmt.Println("Link at i is : ", data[i])
	}
}

/*
Crawl executes the crawling of a website with number of concurrent process assigned
*/
func Crawl(startURL string, parser Parser, concurrency int) []ScrapeResult {
	results := []ScrapeResult{}
	// create a channel worklist that returns a string
	worklist := make(chan []string)
	var n int
	n++
	//this is another channel - but unsure as to its usage
	var tokens = make(chan struct{}, concurrency)
	//go param denotes execute this process asynchronously
	go func() { worklist <- []string{startURL} }()

	// initialise a new map (key is a string and value is a bool)
	seen := make(map[string]bool)
	baseDomain := parseStartURL(startURL)

	for ; n > 0; n-- {
		//receive from worklist channel and assign to list
		list := <-worklist
		for _, link := range list {
			if !seen[link] {
				seen[link] = true
				n++
				go func(baseDomain, link string, parser Parser, token chan struct{}) {
					foundLinks, pageResults := crawlPage(baseDomain, link, parser, token)
					//fmt.Println("found a total of ", len(foundLinks), " links for url: ", link)
					//readAddPrintAllUrls(foundLinks)
					results = append(results, pageResults)
					if foundLinks != nil {
						if len(foundLinks) > 9 {
							fmt.Println("> 9", len(foundLinks), " links for url: ", link)
							worklist <- foundLinks[0:8]
						} else {
							fmt.Println("found a total of ", len(foundLinks), " links for url: ", link)
							worklist <- foundLinks
						}
						//worklist <- foundLinks
					}
				}(baseDomain, link, parser, tokens)
			}
		}
	}
	return results
}
