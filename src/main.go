package main

import (
	"bufio"
	"bytes"
	"fmt"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/epub/manifests"
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

	// npages := make(map[string][]byte)
	{
		pages := fetch.Scrapper(args.Targets[0], args)
		for key, val := range pages {
			fmt.Printf("Page: %s\n", key.String())
			fmt.Printf("Data: %+v\n", val.MediaType)
			// Pay attention to the extra `/` at the start of path
			// stripped := strings.TrimLeft(key.Path, "/")
			// npages[stripped] = val.Data
		}
		buf := make([]byte, 1024*3)
		out := bytes.NewBuffer(buf)
		writeBuffer := bufio.NewWriter(out)
		manifests.GenerateContentOpf(writeBuffer, pages)
		println(string(out.Bytes()))
	}
	// zip.WriteTozip(npages, args.Output+".zip")
}
