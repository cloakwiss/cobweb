package process

import (
	"log"
	"mime"
	"strings"

	"github.com/cloakwiss/cobweb/fetch"
	"github.com/cloakwiss/cobweb/tidy"
)

type AllAssets struct {
	XhtmlPages, Asssets []string
	AllAssetStore       map[string]fetch.Asset
}

func ConvertHtml(allAssets fetch.PageTable) AllAssets {
	// The Uri not always end in .html
	var (
		pageNumber, assetNumber uint
		pages                   = make([]string, len(allAssets))
		assets                  = make([]string, len(allAssets))
		allAssetsStore          = make(map[string]fetch.Asset)
		xhtmlMime               = mime.TypeByExtension(".xhtml")
	)
	for uri, data := range allAssets {
		path := uri.EscapedPath()
		isHtml, newName := toConvert(path)
		// If it is html then insert it in new map
		// for pages
		if isHtml {
			xhtml := tidy.TidyHTML(data.Data)
			pages[pageNumber] = "page:" + newName
			allAssetsStore[pages[pageNumber]] = fetch.Asset{
				Data: xhtml,
				Metadata: fetch.Metadata{
					MediaType: xhtmlMime,
					Title:     data.Title,
				},
			}
			pageNumber += 1
		} else {
			assets[assetNumber] = "asset:" + path
			allAssetsStore[assets[assetNumber]] = fetch.Asset{
				Data: data.Data,
				Metadata: fetch.Metadata{
					MediaType: data.MediaType,
					Title:     data.Title,
				},
			}
			assetNumber += 1
		}
	}
	return AllAssets{pages, assets, allAssetsStore}
}

// This function makes decision to run the tidy html function or not
// and also suggests new name
func toConvert(path string) (bool, string) {
	if strings.HasSuffix(path, "html") {
		newName, found := strings.CutSuffix(path, "html")
		if found {
			newName += "xhtml"
			return true, newName
		} else {
			log.Fatal("Cannot find the html in file's name")
		}
	}
	return false, ""
}

// func GetMetadata(pages fetch.PageTable) {
// 	for url, data := range pages {
// 		url
// 	}
// }
