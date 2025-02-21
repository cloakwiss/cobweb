package fetch_test

import (
	"testing"

	"github.com/cloakwiss/cobweb/fetch"
)

// I have hosted website with caddy on port 8080
// and ran test against it to see if it making map of url and its content

const Target = "http://localhost:8080"

func Test_Mainloop(t *testing.T) {
	fetch.Mainloop(Target, 1)
}
