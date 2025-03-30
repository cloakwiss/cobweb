package main

import (
	"bytes"
	"fmt"
	"log"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/fetch"
	"golang.org/x/net/html"
)

func main() {
	args := app.Args()
	// fmt.Printf("%+v", args)

	// npages := make(map[string][]byte)
	{
		pages := fetch.Scrapper(args.Targets, args)
		for key, val := range pages {
			fmt.Printf("Page: %s\n", key.String())
			fmt.Printf("MediaType: %+v\n", val.MediaType)
			// Pay attention to the extra `/` at the start of path
			// stripped := strings.TrimLeft(key.Path, "/")
			// npages[stripped] = val.Data
		}
		data, found := pages[args.Targets]
		if !found {
			log.Fatalln("It did not download the MAIN PAGE")
		}
		buf := bytes.NewBuffer(data.Data)
		getMetaData(buf)
		// getMetaData(*buf)
		// buf := make([]byte, 1024*3)
		// out := bytes.NewBuffer(buf)
		// writeBuffer := bufio.NewWriter(out)
		// manifests.GenerateContentOpf(writeBuffer, pages)
		// println(string(out.Bytes()))
	}
	// zip.WriteTozip(npages, args.Output+".zip")
}

type MetaData struct {
	title string
	other map[string][]string
}

func getMetaData(buf *bytes.Buffer) MetaData {
	doc, err := html.Parse(buf)
	if err != nil {
		log.Fatalln("Error parsing HTML:", err)
	}

	head := findHead(doc)
	if head == nil {
		log.Fatalln("Cannot find head.")
	}
	return processHeadElements(head)
}

// Function to find the head element
func findHead(n *html.Node) *html.Node {
	if n.Type == html.ElementNode && n.Data == "head" {
		return n
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		if result := findHead(c); result != nil {
			return result
		}
	}

	return nil
}

// Function to process elements inside head
func processHeadElements(head *html.Node) MetaData {
	data := MetaData{
		other: make(map[string][]string),
	}

	for child := head.FirstChild; child != nil; child = child.NextSibling {
		if child.Type == html.ElementNode {

			// // Print attributes if any
			if len(child.Attr) > 0 {
				for _, attr := range child.Attr {
					_, found := data.other[attr.Key]
					if found {
						data.other[attr.Key] = append(data.other[attr.Key], attr.Val)
					} else {
						data.other[attr.Key] = []string{attr.Val}
					}
				}
			}

			// If it's a title element, print its content
			if child.Data == "title" && child.FirstChild != nil {
				data.title = child.FirstChild.Data
			}
		}
	}
	return data
}
