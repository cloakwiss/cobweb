/*
This modules handles all the crawler and also rewriting of the body which means it will interact with
*Html Sanitizer* and *User Interface*

Output of this will be SiteMap (Graph of all the pages hopefully Directed Acyclic Graph),
Body of all the Pages and continous update for UI through a channel
*/
package fetch

import (
	"fmt"
	"github.com/gocolly/colly"
	// "net/url"
)

// Obtained from Brave Browser
const DefaultUserAgent = "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/133.0.0.0 Safari/537.36"
const Target = "https://matklad.github.io/"

// This is the function that will be start the crawling and scrapping
func Start(inchan chan int, outchan chan int) {}

// This is supposed to find all path in
func mainloop() {
	c := colly.NewCollector(
		colly.AllowedDomains("matklad.github.io"),
	)

	c.UserAgent = DefaultUserAgent
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		fmt.Printf("Link found: %q -> %s\n", e.Text, link)
	})

	c.OnResponse(func(res *colly.Response) {
		fmt.Printf("Body of %s:\n %s\n", res.Request.URL, string(res.Body))
	})

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL.String())
	// })

	c.Visit(Target)
}
