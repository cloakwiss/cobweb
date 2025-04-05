package core

import (
	"bufio"
	"bytes"
	"fmt"
	"log"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/epub/manifests"
	"github.com/cloakwiss/cobweb/epub/process"
	"github.com/cloakwiss/cobweb/epub/zip"
	"github.com/cloakwiss/cobweb/fetch"
)

func Launch(args app.Options) string {
	pages := fetch.Scrapper(args.Targets, args)

	for key := range pages {
		fmt.Printf("Page: %s\n\tPath: %s\n", key.String(), key.EscapedPath())

		// fmt.Printf("MediaType: %+v\n", val.MediaType)
		// Pay attention to the extra `/` at the start of path
		// stripped := strings.TrimLeft(key.Path, "/")
		// npages[stripped] = val.Data
	}

	// Process the pages
	processed := process.OrderAndConvertPages(pages)
	// for _, k := range processed.XhtmlPages {
	// 	fmt.Printf("\t`%s`\n", k)
	// 	// fmt.Println(string(propages.AllAssetStore[k].Data))
	// }
	// for _, k := range processed.Assets {
	// 	fmt.Printf("\t`%s`\n", k)
	// 	// fmt.Println(string(propages.AllAssetStore[k].Data))
	// }

	// Fetch metadata from the first page
	var metaData fetch.PageMetadata
	{
		data, found := pages[args.Targets]
		if !found {
			log.Fatalln("It did not download the MAIN PAGE")
		}
		mainPage := bytes.NewBuffer(data.Data)
		metaData = fetch.GetMetaData(mainPage)
	}

	var containerXml []byte = make([]byte, 0, 1024)
	{
		pbuf := make([]byte, 0, 1024)
		out := bytes.NewBuffer(pbuf)
		writeBuffer := bufio.NewWriter(out)
		manifests.NewContainer(writeBuffer, "content.opf")
		containerXml = append(containerXml, out.Bytes()...)
	}
	// println(string(containerXml))

	var contentOpf []byte = make([]byte, 0, 1024*3)
	{
		pbuf := make([]byte, 0, 1024*3)
		out := bytes.NewBuffer(pbuf)
		writeBuffer := bufio.NewWriter(out)
		manifests.GenerateContentOpf(writeBuffer, metaData, processed)
		contentOpf = append(contentOpf, out.Bytes()...)
	}
	// println(string(contentOpf.Bytes()))

	var toc []byte = make([]byte, 0, 1024*3)
	{
		pbuf := make([]byte, 0, 1024*3)
		out := bytes.NewBuffer(pbuf)
		writeBuffer := bufio.NewWriter(out)
		for i := range processed.XhtmlPages {
			println(processed.XhtmlPages[i])
		}
		manifests.MarshalToc(manifests.GenerateDirectoryTree(processed.XhtmlPages), writeBuffer)
		toc = append(toc, out.Bytes()...)
	}
	// println(string(toc.Bytes()))
	filesBlobs := make([]zip.Pair, 0, len(processed.AllAssetStore)+3)
	filesBlobs = append(filesBlobs, zip.Pair{File: "META-INF/container.xml", Bytes: containerXml})
	filesBlobs = append(filesBlobs, zip.Pair{File: "content.opf", Bytes: contentOpf})
	filesBlobs = append(filesBlobs, zip.Pair{File: "nav.xhtml", Bytes: toc})

	for _, filename := range processed.XhtmlPages {
		blob, found := processed.AllAssetStore[filename]
		if !found {
			log.Printf("Not found: `%s`\t `%+v`", filename, blob.Metadata)
		} else {
			filesBlobs = append(filesBlobs, zip.Pair{
				File:  filename,
				Bytes: blob.Data,
			})
		}
	}
	for _, filename := range processed.Assets {
		blob, found := processed.AllAssetStore[filename]
		if !found {
			log.Printf("Not found: `%s`", filename)
		} else {
			filesBlobs = append(filesBlobs, zip.Pair{
				File:  filename,
				Bytes: blob.Data,
			})
		}
	}
	zip.WriteTozip(filesBlobs, args.Output+".epub")

	return args.Output + ".epub"
}
