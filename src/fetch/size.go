package fetch

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func handleHeaders(resp *colly.Response) (uint64, bool) {
	var size uint64 = 0
	var getworthy bool = false

	fmt.Printf("Headers: {\n[%v]\n}\n", resp.Headers)

	// Getting size from header
	length := resp.Headers.Get("Content-Length")
	urlsize, err := strconv.Atoi(length)
	if err != nil {
		fmt.Printf("Error getting head for url [%s]: err [%v]\n", resp.Request.URL.String(), err)
	}
	size += uint64(urlsize)

	// Content-type check
	content_type := strings.Split(resp.Headers.Get("Content-Type"), ";")[0]
	fmt.Printf("Content-type: %s\n", content_type)

	// TODO: Need to check bunch of http content_type tags
	getworthy = strings.Contains(content_type, "html")

	return size, getworthy
}

func ScrapeSize(target string, recurse_limit uint8) uint64 {
	return handleGetworthy(target, recurse_limit)
}

func handleGetworthy(target string, recurse_limit uint8) uint64 {
	var size uint64 = 0

	targetUrl, err := url.Parse(target)
	if err != nil {
		fmt.Printf("Error getting url [%s]: err [%v]\n", target, err)
	}

	collector := colly.NewCollector(
		colly.AllowedDomains(removeProtocolPrefix(targetUrl)),
		colly.MaxDepth(int(recurse_limit)+1),
	)

	// Need to add others too
	{
		htmlHandler := func(e *colly.HTMLElement) {
			link := e.Attr("href")
			linksize, worthy := ScrapeSizeAux(link)
			if worthy {
				e.Request.Visit(link)
			}
			size += linksize
		}

		//TODO: verify if this is all
		collector.OnHTML("a[href]", htmlHandler)
		collector.OnHTML("link[href]", htmlHandler)
		collector.OnHTML("base[href]", htmlHandler)
		collector.OnHTML("area[href]", htmlHandler)
		collector.OnHTML("data[object]", htmlHandler)
		collector.OnHTML("del[cite]", htmlHandler)
		collector.OnHTML("ins[cite]", htmlHandler)
		collector.OnHTML("blockquote[cite]", htmlHandler)
		collector.OnHTML("q[cite]", htmlHandler)
		collector.OnHTML("img[src]", htmlHandler)
		collector.OnHTML("track[src]", htmlHandler)
		collector.OnHTML("embed[src]", htmlHandler)
		collector.OnHTML("source[src]", htmlHandler)
		collector.OnHTML("script[src]", htmlHandler)
		collector.OnHTML("iframe[src]", htmlHandler)
	}

	targetsize, worthy := ScrapeSizeAux(target)
	if worthy {
		collector.Visit(target)
	}

	size += targetsize

	return size
}

func ScrapeSizeAux(target string) (uint64, bool) {
	var size uint64 = 0

	targetUrl, err := url.Parse(target)
	if err != nil {
		fmt.Printf("Error getting url [%s]: err [%v]\n", target, err)
	}
	collector := colly.NewCollector(
		colly.AllowedDomains(removeProtocolPrefix(targetUrl)),
		colly.MaxDepth(1),
	)

	// fmt.Printf("Scraping sizes.\n")

	getworthy := false
	collector.OnResponse(func(resp *colly.Response) {
		println("heading")
		targetsize, recursable := handleHeaders(resp)
		println("head ended")

		getworthy = recursable
		size += targetsize
	})

	// Head Request
	head_error := collector.Head(target)
	if err != nil {
		fmt.Printf("Error HEAD requesting for url [%s]: err [%v]\n", target, head_error)
	}

	return size, getworthy
}
