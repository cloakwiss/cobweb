package fetch

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/cloakwiss/cobweb/app"
	"github.com/gocolly/colly"
)

// These are header collected from some pages have not figured out how to use it effectively
// Ideally these should not be required
var header = map[string][]string{
	// Obtained from Firefox Browser
	"Accept-Encoding": {"gzip", "deflate", "br", "zstd"},
	"Accept":          {"text/html"},
	"User-Agent": {
		// Obtained from Brave Browser
		"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36",
		// Obtained from Firefox Browser
		"Mozilla/5.0 (X11; Linux x86_64; rv:135.0) Gecko/20100101 Firefox/135.0",
	},
}

// This function should be called for all the target URL seprately
// or we need to find out if we can safely merge the graph of page
// discovered without redoing the work.
//
// For example if the urls are pointing to same domain but are of different
// path then we can share the downloaded result to some extend atleast.
//
//	And if that is the case then the index (of the book) generation will also need to change
//
// TODO: added some headers, caching and URL filter smarter
// TODO: there is a feature to use regex to mactch urls, need to find out how to expose it to users
func Scrapper(target url.URL, argu app.Options, out chan<- app.ApMsg) map[url.URL][]byte {
	// defer close(out)
	var pagesContents map[url.URL][]byte = make(map[url.URL][]byte)
	println("Domain Name: ", target.String())
	println("Depth: ", argu.Depth)

	// recurse limit is unused
	collector := colly.NewCollector(
		colly.AllowedDomains(stringOfURL(argu.AllowDomains)...),
		colly.DisallowedDomains(stringOfURL(argu.BlockDomains)...),
		colly.MaxDepth(int(argu.Depth)+1),
		// Need to measure the difference in performance (using some older version)
		// colly.Async(true),
	)

	//TODO: Need to add a delay too and expose as flag to cli
	// c.Limit(&colly.LimitRule{
	// 	Delay: time.Second,
	// })

	//TODO: do not change the order of these callback methods
	// https://go-colly.org/docs/introduction/start/ read this
	//-------------------------------------------------------

	collector.OnRequest(func(r *colly.Request) {
		// r.Headers = (*http.Header)(&header)
		// out <- app.ApMsg{
		// 	Code: app.VisitingPage,
		// 	URL:  r.URL.String(),
		// }
		fmt.Println("Visiting", r.URL.String())
	})

	collector.OnError(func(r *colly.Response, err error) {})

	collector.OnResponse(func(res *colly.Response) {
		pagesContents[*res.Request.URL] = res.Body
		// out <- app.ApMsg{
		// 	Code: app.OnPage,
		// 	URL:  res.Request.URL.String(),
		// }
		fmt.Printf("On page: %v\n", res.Request.URL)
	})

	// Need to add others too
	{
		// TODO: need to find out how can I apply the filter in this function for various resources
		htmlHandler := func(tag string) colly.HTMLCallback {
			return func(e *colly.HTMLElement) {
				link := e.Attr(tag)
				e.Request.Visit(link)
			}
		}

		// TODO: verify if this is all
		collector.OnHTML("a[href]", htmlHandler("href"))
		collector.OnHTML("link[href]", htmlHandler("href"))
		collector.OnHTML("base[href]", htmlHandler("href"))
		collector.OnHTML("area[href]", htmlHandler("href"))

		collector.OnHTML("data[object]", htmlHandler("object"))

		collector.OnHTML("del[cite]", htmlHandler("cite"))
		collector.OnHTML("ins[cite]", htmlHandler("cite"))
		collector.OnHTML("blockquote[cite]", htmlHandler("cite"))
		collector.OnHTML("q[cite]", htmlHandler("cite"))

		collector.OnHTML("img[src]", htmlHandler("src"))
		collector.OnHTML("track[src]", htmlHandler("src"))
		collector.OnHTML("embed[src]", htmlHandler("src"))
		collector.OnHTML("source[src]", htmlHandler("src"))
		collector.OnHTML("script[src]", htmlHandler("src"))
		collector.OnHTML("iframe[src]", htmlHandler("src"))
	}

	collector.OnScraped(func(r *colly.Response) {})

	println("Started the scrapper")
	collector.Visit(target.String())

	//-------------------------------------------------------
	return pagesContents
}

// Stick to url struct as much as possible
func removeProtocolPrefix(url *url.URL) string {
	var hostname string
	if url.Port() == "" {
		hostname = url.Hostname()
	} else {
		hostname = url.Hostname() + ":" + url.Port()
	}
	return strings.TrimLeft(hostname, "www.")
}

// WARN: this cannot handle paths at the moment
// TODO: what should the behaviour of this in case the it has a username in the url
func stringOfURL(urls []url.URL) []string {
	parsed := make([]string, 0, len(urls))
	for _, u := range urls {
		// u.String()
		if u.Scheme == "http" || u.Scheme == "https" {
			//WARN: This url's port is stripped here
			ur := u.Hostname()
			parsed = append(parsed, ur)
		}
	}
	return parsed
}
