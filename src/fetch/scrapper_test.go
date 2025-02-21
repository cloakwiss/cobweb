package fetch_test

import (
	"testing"

	"github.com/cloakwiss/cobweb/fetch"
)

const Target = "http://localhost:8080"

func Test_Mainloop(t *testing.T) {
	fetch.Mainloop(Target, 1)
}
