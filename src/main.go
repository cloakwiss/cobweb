package main

import (
	"fmt"
	"strings"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/epub/zip"
	"github.com/cloakwiss/cobweb/fetch"
)

func main() {
	args := app.Args()
	// fmt.Printf("%+v", args)

	if len(args.Targets) == 0 {
		println("Early exit")
		return
	}
	fmt.Printf("Targets: %+v\n", args.Targets)

	npages := make(map[string][]byte)
	{
		pages := fetch.Scrapper(args.Targets[0], args)
		for key, val := range pages {
			// fmt.Printf("Page: %s\n", key.String())
			// Pay attention to the extra `/` at the start of path
			stripped := strings.TrimLeft(key.Path, "/")
			npages[stripped] = val.Data
		}
	}
	zip.WriteTozip(npages, args.Output+".zip")
}
