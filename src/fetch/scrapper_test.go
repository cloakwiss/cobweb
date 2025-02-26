package fetch_test

import (
	"fmt"
	"testing"

	"github.com/cloakwiss/cobweb/fetch"
)

// This is rust's documentation which is hosted by caddy to run test
const target = "http://localhost:8080/std/index.html"

func Test_Mainloop(t *testing.T) {
	// pages := fetch.Mainloop(target, 1)
	// for key := range pages {
	// 	fmt.Printf("Page: %s\n", key.String())
	// }

	size := fetch.ScrapeSize(target, 1)
	fmt.Printf("total size: %d \n", size)
}
