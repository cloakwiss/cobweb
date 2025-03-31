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
		pageNumber, assetNumber uint
		pages                   = make([]string, len(allAssets))
		assets                  = make([]string, len(allAssets))
		allAssetsStore          = make(map[string]fetch.Asset)
		xhtmlMime               = mime.TypeByExtension(".xhtml")
		htmlMime                = mime.TypeByExtension(".html") // Also pay attention to encoding
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
		path := uri.EscapedPath()
		// If it is html then insert it in new map
		// for pages
		if strings.Contains(data.MediaType, htmlMime) {
			xhtml := tidy.TidyHTML(data.Data)
			pages[pageNumber] = newName(path)
			allAssetsStore[pages[pageNumber]] = fetch.Asset{
				Data: xhtml,
				Metadata: fetch.Metadata{
					MediaType: xhtmlMime,
				},
			}
			pageNumber += 1
		} else {
			assets[assetNumber] = path
			allAssetsStore[assets[assetNumber]] = fetch.Asset{
				Data: data.Data,
				Metadata: fetch.Metadata{
					MediaType: data.MediaType,
				},
			}
			assetNumber += 1
		}
	}
	return AllAssets{pages, assets, allAssetsStore}
}

func newName(path string) string {
	if strings.HasSuffix(path, "html") {
		newName, found := strings.CutSuffix(path, "html")
		if found {
			newName += "xhtml"
			return newName
		} else {
			log.Fatal("Unreachable")
		}
	}
	return path + ".xhtml"
}
