package process

import (
	"log"
	"mime"
	"net/url"
	"slices"
	"strings"

	"github.com/cloakwiss/cobweb/fetch"
	"github.com/cloakwiss/cobweb/tidy"
)

type AllAssets struct {
	XhtmlPages, Assets []string
	AllAssetStore      map[string]fetch.Asset
}

func OrderAndConvertPages(allAssets fetch.PageTable) AllAssets {
	// The Uri not always end in .html
	var (
		pageNumber, assetNumber uint = 0, 0
		pages                        = make([]string, len(allAssets))
		assets                       = make([]string, len(allAssets))
		allAssetsStore               = make(map[string]fetch.Asset)
		xhtmlMime                    = mime.TypeByExtension(".xhtml")
		htmlMime                     = "text/html" // Also pay attention to encoding
	)
	keys := make([]url.URL, 0, len(allAssets))
	for u := range allAssets {
		keys = append(keys, u)
	}
	slices.SortFunc(keys, func(a, b url.URL) int {
		return strings.Compare(a.EscapedPath(), b.EscapedPath())
	})

	for _, uri := range keys {
		data := allAssets[uri]
		path := strings.TrimPrefix(uri.EscapedPath(), "/")
		// minor hack
		if path == "" {
			if strings.Contains(data.MediaType, htmlMime) {
				log.Println("Edge case path of ``")
				path = "root.html"
			}
		} else if path == "/" {
			log.Println("Edge case path of `/`: ", path)
		} else if strings.HasSuffix(path, "/") {
			path = strings.TrimSuffix(path, "/")
			path += ".html"
		}

		if strings.Contains(data.MediaType, htmlMime) {
			xhtml := tidy.TidyHTML(data.Data)
			if xhtml == nil {
				log.Printf("Path: %s", path)
				continue
			}
			println("Path: ", path, "Length: ", len(xhtml))
			newpath := newName(path)
			pages[pageNumber] = newpath
			allAssetsStore[newpath] = fetch.Asset{
				Data: xhtml,
				Metadata: fetch.Metadata{
					MediaType: xhtmlMime,
				},
			}
			pageNumber += 1
		} else {
			assets[assetNumber] = path
			allAssetsStore[path] = fetch.Asset{
				Data: data.Data,
				Metadata: fetch.Metadata{
					MediaType: data.MediaType,
				},
			}
			assetNumber += 1
		}
	}
	return AllAssets{
		XhtmlPages:    slices.Compact(pages),
		Assets:        slices.Compact(assets),
		AllAssetStore: allAssetsStore,
	}
}

func newName(path string) string {
	if strings.HasSuffix(path, ".html") {
		newName, found := strings.CutSuffix(path, ".html")
		if found {
			newName += ".xhtml"
			return newName
		} else {
			log.Fatal("Unreachable")
		}
	}
	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
		return path + ".xhtml"
	}
	if path == "" {
		return "base.xhtml"
	}
	path += ".xhtml"
	return path
}
