package fetch_test

import (
	"fmt"
	"testing"

	"github.com/cloakwiss/cobweb/fetch"
)

// This is rust's documentation which is hosted by caddy to run test
const target = "http://localhost:8080/std/index.html"

func Test_Mainloop(t *testing.T) {
	pages := fetch.Mainloop(target, 1)
	npages := make(map[string][]byte)
	for key, val := range pages {
		fmt.Printf("Page: %s\n", key.String())
		npages[key.Path] = val
	}
	fetch.WriteTozip(npages, "output.zip")
}
