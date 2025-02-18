package main

import (
	xml "github.com/cloakwiss/cobweb/xml"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		xml.Start(arg)
	}
}
