/*
This modules handles all the crawler and also rewriting of the body which means it will interact with
*Html Sanitizer* and *User Interface*

Output of this will be SiteMap (Graph of all the pages hopefully Directed Acyclic Graph),
Body of all the Pages and continous update for UI through a channel
*/
package fetch

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/gocolly/colly"
	// "net/url"
)

// Obtained from Brave Browser
const DefaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"

// This is the function that will be start the crawling and scrapping
// func Start(inchan chan int, outchan chan int) {}

// This is supposed to find all path in the page with some limit on recursion
func Mainloop(target string, recurse_limit uint8) map[url.URL][]byte {
	// recurse limit is unused
	_ = recurse_limit
	c := colly.NewCollector(
		colly.AllowedDomains(removeProtocolPrefix(target)),
	)
	c.UserAgent = DefaultUserAgent

	var pages map[url.URL][]byte = make(map[url.URL][]byte)

	// Need to add others too
	{
		// Anchor will not be the only tag used to link to other pages
		c.OnHTML("a[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(e.Request.AbsoluteURL(link))
		})

		// will look for other assets
		c.OnHTML("link[href]", func(e *colly.HTMLElement) {
			link := e.Attr("href")
			c.Visit(e.Request.AbsoluteURL(link))
		})
	}

	c.OnResponse(func(res *colly.Response) {
		pages[*res.Request.URL] = res.Body
		fmt.Printf("On page: %v\n", res.Request.URL)
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	c.Visit(target)

	for key := range pages {
		fmt.Printf("Page: %s\n", key.EscapedPath())
	}

	return pages
}

func removeProtocolPrefix(url string) string {
	_, after, found := strings.Cut(url, "//")
	if found {
		return after
	} else {
		return ""
	}
}
