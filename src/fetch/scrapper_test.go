package fetch_test

import (
	"testing"

	"github.com/cloakwiss/cobweb/fetch"
)

const Target = "http://stract.com/"

func Test_Mainloop(t *testing.T) {
	fetch.Mainloop(Target, 1)
}
