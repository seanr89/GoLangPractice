package crawler

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
	"github.com/PuerkitoBio/goquery"
)

/*
Here’s the worker, of which we’ll run several concurrent instances.
These workers will receive work on the jobs channel and send the corresponding results on results.
*/
func worker(id int, baseURL string, jobs <-chan string, result chan<- string) {
	fmt.Println("worker execution: ", id)
	if jobs != nil {
		for j := range jobs {
			res := crawlCurrentPage(baseURL, j, id)
			if res != nil{
				//result <- res
				//return result
				for _, v := range res{
					//fmt.Println("result item", v)
					result <- v
				}
			}
		}
	} else {
		fmt.Println("jobs was nil")
	}
}

func outputURLS(urls []string) {
	if urls != nil {
		fmt.Println("length of urls =", len(urls))
		for _, href := range urls {
			fmt.Println("url = : ", href)
		}
	}
}

/*
 */
func resolveRelative(baseURL string, hrefs []string) []string {
	//fmt.Println("resolveRelative: ", hrefs, "with base: ", baseURL)
	internalUrls := []string{}

	for _, href := range hrefs {
		//fmt.Println("parsing href: ", href)
		if strings.HasPrefix(href, baseURL) {
			internalUrls = append(internalUrls, href)
		}

		///First iteration
		//if strings.HasPrefix(href, "/") {
			// resolvedURL := fmt.Sprintf("%s", href)
			// fmt.Println("resolvedURL href: ", resolvedURL)
			// internalUrls = append(internalUrls, resolvedURL)
		//}


		/// another methodbeing testing
		if strings.HasPrefix(href, "/") {
			resolvedURL := fmt.Sprintf("%s%s", baseURL, href)
			//fmt.Println("resolvedURL href: ", resolvedURL)
			internalUrls = append(internalUrls, resolvedURL)
		}
	}
	return internalUrls
}

/*
 */
func crawlCurrentPage(baseURL, targetURL string, id int) []string {
	fmt.Println("crawlCurrentPage: ", targetURL, "worker id: ", id)
	resp, _ := getURLRequest(targetURL)

	doc, _ := goquery.NewDocumentFromResponse(resp)
	//pageResults := parser.ParsePage(doc)
	links := extractURLLinks(doc)
	if links != nil{
	foundUrls := resolveRelative(baseURL, links)
	return foundUrls
	}
	return nil
}

/*
parseStartURL output the scheme and the host of the site
returns an output string to console with the scheme and host
*/
func parseStartingURL(u string) string {
	fmt.Println("parseStartingURL")
	// use url package to parse the url from the string
	parsed, _ := url.Parse(u)
	return fmt.Sprintf("%s://%s", parsed.Scheme, parsed.Host)
}

/*
 */
func getURLRequest(url string) (*http.Response, error) {
	//fmt.Println("getURLRequest: ", url)
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
func extractURLLinks(doc *goquery.Document) []string {
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

func StartCrawl(baseURL string, workerCount int) {
	fmt.Println("StartCrawl: ", baseURL, workerCount)

	resultURL := make(chan string, 10000)

	// create a channel worklist that expects a string
	//the worklist will be an array or urls (appended to during crawling)
	workList := make(chan string)
	//baseURLChan := make(chan string, 200)

	// loop through and and create workers - initially stopped as no jobs(worklist) present
	for w := 1; w <= workerCount; w++ {
		//fmt.Println("for worker")
		go worker(w, baseURL, workList, resultURL)
	}

	var parsedStartingURL = parseStartingURL(baseURL)
	workList <- parsedStartingURL
	
	for a := 1; a <= 1000; a++ {
		time.Sleep(1 * time.Second)
		res := <- resultURL
		//fmt.Println("for res:", res)
		workList <- res
	}
	fmt.Println("complete:")

	//time to now handle any response from each worker iteration
	fmt.Scanln()
	close(workList)
}
