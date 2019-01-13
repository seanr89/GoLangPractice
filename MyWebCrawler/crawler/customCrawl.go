package crawler

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

/*
Here’s the worker, of which we’ll run several concurrent instances.
These workers will receive work on the jobs channel and send the corresponding results on results.
*/
func worker(id int, jobs <-chan string, results chan<- []string) {
	fmt.Println("worker execution")
	if jobs != nil {
		for j := range jobs {
			results := crawlCurrentPage(j)
		}
	}
}

func crawlCurrentPage(targetURL string) []string {
	fmt.Println("Requesting: ", targetURL)
	resp, _ := getURLRequest(targetURL)

	doc, _ := goquery.NewDocumentFromResponse(resp)
	//pageResults := parser.ParsePage(doc)
	links := extractLinks(doc)
	foundUrls := make([]string, 0)
	//foundUrls := resolveRelative(baseURL, links)

	return foundUrls
}

/*
parseStartURL output the scheme and the host of the site
returns an output string to console with the scheme and host
*/
func parseStartingURL(u string) string {
	// use url package to parse the url from the string
	parsed, _ := url.Parse(u)
	return fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)
}

func getURLRequest(url string) (*http.Response, error) {
	fmt.Println("getURLRequest: ", url)
	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)")

	res, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	return res, nil
}

func StartCrawl(baseURL string, workerCount int) {
	//initialise array of scraperesult objects taken from the controller
	results := []ScrapeResult{}

	// create a channel worklist that returns a string
	//the worklist will be an array or urls (appended to during crawling)
	workList := make(chan string)
	resultChan := make(chan []string)

	var parsedStartingURL = parseStartingURL(baseURL)
	workList <- parsedStartingURL

	// loop through and and create workers - initially stopped as no jobs present
	for w := 1; w <= workerCount; w++ {
		go worker(w, workList, resultChan)
	}

	//time to now handle any response from each worker iteration
}
