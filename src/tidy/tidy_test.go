package tidy_test

import (
	"fmt"
	"testing"

	"github.com/cloakwiss/cobweb/tidy"
)

// This test won't run unless the dll/shared object for tidy is in the same
// folder as this package
func TestTidyHTML(t *testing.T) {
	input := []byte("<html><head><title>Test</title></head><body><p>Test</p></body></html>")
	output := tidy.TidyHTML(input)

	fmt.Println(string(output))
	if len(output) == 0 {
		t.Errorf("TidyHTML returned an empty result")
	}
}
