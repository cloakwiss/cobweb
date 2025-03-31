package main

import (
	"fmt"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/epub/process"
	"github.com/cloakwiss/cobweb/fetch"
)

func main() {
	args := app.Args()
	// fmt.Printf("%+v", args)

	pages := fetch.Scrapper(args.Targets, args)

	for key := range pages {
		fmt.Printf("Page: %s\n", key.String())
		// fmt.Printf("MediaType: %+v\n", val.MediaType)
		// Pay attention to the extra `/` at the start of path
		// stripped := strings.TrimLeft(key.Path, "/")
		// npages[stripped] = val.Data
	}

	// Process the pages
	propages := process.OrderAndConvertPages(pages)
	for _, k := range propages.XhtmlPages {
		fmt.Println("\t`", k, "`")
		// fmt.Println(string(propages.AllAssetStore[k].Data))
	}
	for _, k := range propages.Assets {
		fmt.Println("\t`", k, "`")
		// fmt.Println(string(propages.AllAssetStore[k].Data))
	}
	// for _, k := range propages.Assets {
	// 	fmt.Println("\t", k)
	// 	fmt.Println(string(propages.AllAssetStore[k].Data))
	// }

	// Fetch metadata from the first page
	// var metaData fetch.PageMetadata
	// {
	// 	data, found := pages[args.Targets]
	// 	if !found {
	// 		log.Fatalln("It did not download the MAIN PAGE")
	// 	}
	// 	mainPage := bytes.NewBuffer(data.Data)
	// 	metaData = fetch.GetMetaData(mainPage)
	// }

	// var contentOpf bytes.Buffer
	// {
	// 	pbuf := make([]byte, 1024*3)
	// 	out := bytes.NewBuffer(pbuf)
	// 	writeBuffer := bufio.NewWriter(out)
	// 	manifests.GenerateContentOpf(writeBuffer, metaData, pages)
	// 	contentOpf = *out
	// }

	// var contentOpf bytes.Buffer
	// {
	// 	pbuf := make([]byte, 1024*3)
	// 	out := bytes.NewBuffer(pbuf)
	// 	writeBuffer := bufio.NewWriter(out)
	// 	manifests.GenerateContentOpf(writeBuffer, metaData, pages)
	// 	contentOpf = *out
	// }
}
