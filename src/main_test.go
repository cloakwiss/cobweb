package main

import (
	"fmt"
	"testing"

	"github.com/cloakwiss/cobweb/epub/zip"
	"github.com/cloakwiss/cobweb/fetch"
)

// This is rust's documentation which is hosted by caddy to run test
const target = "http://localhost:8080/std/index.html"

func Test_Mainloop(t *testing.T) {
	pages := fetch.Mainloop(target, 1)
	npages := make(map[string][]byte)
	for key, val := range pages {
		fmt.Printf("Page: %s\n", key.String())
		// Pay attention to the extra `/` at the start of path
		npages[key.Path] = val
	}
	zip.WriteTozip(npages, "output.zip")
}
