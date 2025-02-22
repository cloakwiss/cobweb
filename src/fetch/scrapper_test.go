package fetch_test

import (
	"testing"

	"github.com/cloakwiss/cobweb/fetch"
)

// This is rust's documentation which is hosted by caddy to run test
const target = "http://localhost:8088/std/index.html"

func Test_Mainloop(t *testing.T) {
	fetch.Mainloop(target, 1)
}
