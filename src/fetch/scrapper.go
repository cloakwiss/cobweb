package fetch

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
)

var header = map[string][]string{
	// Obtained from Firefox Browser
	"Accept-Encoding": {"gzip", "deflate", "br", "zstd"},
	"Accept":          {"text/html", "application/xhtml+xml", "application/xml;q=0.9,*/*;q=0.8"},
	// Obtained from Brave Browser
	"User-Agent": {"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"},
}

// This is the function that will be start the crawling and scrapping
// func Start(inchan chan int, outchan chan int) {}

// This is supposed to find all path in the page with some limit on recursion
// TODO: added some headers and caching also make the URL filter smarter
func Mainloop(target string, recurse_limit uint8) map[url.URL][]byte {
	targetUrl, err := url.Parse(target)
	var pagesContents map[url.URL][]byte = make(map[url.URL][]byte)
	if err != nil {
		return pagesContents
	}
	println("Domain Name: ", removeProtocolPrefix(targetUrl))

	// recurse limit is unused
	c := colly.NewCollector(
		colly.AllowedDomains(removeProtocolPrefix(targetUrl)),
		colly.MaxDepth(int(recurse_limit)),
	)

	//TODO: Need to add a delay too
	// c.Limit(&colly.LimitRule{
	// 	Delay: time.Second,
	// })

	//TODO: do not change the order of these callback methods
	// https://go-colly.org/docs/introduction/start/ read this
	//-------------------------------------------------------

	c.OnRequest(func(r *colly.Request) {
		r.Headers = (*http.Header)(&header)
		fmt.Println("Visiting", r.URL.String())
	})

	c.OnError(func(r *colly.Response, err error) {})

	c.OnResponse(func(res *colly.Response) {
		pagesContents[*res.Request.URL] = res.Body
		fmt.Printf("On page: %v\n", res.Request.URL)
	})

	// Need to add others too
	{
		htmlHandler := func(e *colly.HTMLElement) {
			link := e.Attr("href")
			e.Request.Visit(link)
		}

		// Anchor will not be the only tag used to link to other pages
		c.OnHTML("a[href]", htmlHandler)
		// will look for other assets
		c.OnHTML("link[href]", htmlHandler)
	}

	c.OnScraped(func(r *colly.Response) {
	})

	c.Visit(target)

	//-------------------------------------------------------

	for key := range pagesContents {
		fmt.Printf("Page: %s\n", key.String())
	}

	return pagesContents
}

func removeProtocolPrefix(url *url.URL) string {
	var hostname string
	if url.Port() == "" {
		hostname = url.Hostname()
	} else {
		hostname = url.Hostname() + ":" + url.Port()
	}
	return strings.TrimLeft(hostname, "www.")
}
