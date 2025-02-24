package fetch

import (
	"archive/zip"
	"bytes"
	"log"
	"net/http"
	"net/url"
	"os"
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
	collector := colly.NewCollector(
		colly.AllowedDomains(removeProtocolPrefix(targetUrl)),
		colly.MaxDepth(int(recurse_limit)+1),
	)

	//TODO: Need to add a delay too
	// c.Limit(&colly.LimitRule{
	// 	Delay: time.Second,
	// })

	//TODO: do not change the order of these callback methods
	// https://go-colly.org/docs/introduction/start/ read this
	//-------------------------------------------------------

	collector.OnRequest(func(r *colly.Request) {
		r.Headers = (*http.Header)(&header)
		// fmt.Println("Visiting", r.URL.String())
	})

	collector.OnError(func(r *colly.Response, err error) {})

	collector.OnResponse(func(res *colly.Response) {
		pagesContents[*res.Request.URL] = res.Body
		// fmt.Printf("On page: %v\n", res.Request.URL)
	})

	// Need to add others too
	{
		htmlHandler := func(e *colly.HTMLElement) {
			link := e.Attr("href")
			e.Request.Visit(link)
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

	collector.OnScraped(func(r *colly.Response) {
	})

	collector.Visit(target)

	//-------------------------------------------------------

	return pagesContents
}

func WriteTozip(pages map[string][]byte, outputZipFile string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	for path, file := range pages {
		f, err := w.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(file)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	zipfile, err := os.Create(outputZipFile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := zipfile.Write(buf.Bytes()); err != nil {
		log.Fatal(err)
	}
	if err := zipfile.Close(); err != nil {
		log.Fatal(err)
	}
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
