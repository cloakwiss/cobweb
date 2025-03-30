package manifests

import (
	"bufio"
	"encoding/xml"
	"strings"
)

// Container.xml
type container struct {
	// XMLName struct{}  `xml:"container"`
	Root    rootfiles `xml:"rootfiles"`
	Version string    `xml:"version,attr"`
	Xmlns   string    `xml:"xmlns,attr"`
}

type rootfiles struct {
	IRoot rootfile `xml:"rootfile"`
}

type rootfile struct {
	Fullpath  string `xml:"full-path,attr"`
	Mediatype string `xml:"media-type,attr"`
}

func NewContainer(writeBuffer *bufio.Writer, path string) error {
	container := container{
		Version: "1.0",
		// Is this namespace required
		Xmlns: "urn:oasis:names:tc:opendocument:xmlns:container",
		Root: rootfiles{
			IRoot: rootfile{
				Fullpath:  path,
				Mediatype: "application/oebps-package+xml",
			},
		},
	}
	header := []byte(xml.Header)
	_, er := writeBuffer.Write(header)
	if er != nil {
		return er
	}
	by, er := xml.MarshalIndent(container, "", "  ")
	if er != nil {
		return er
	}
	_, er = writeBuffer.Write(by)
	if er != nil {
		return er
	}
	return nil
}

// content.opf

// metadata section of content.opf

// var DublinCoreElements = [][]string{
// 	// Required
// 	{"identifier", "title", "language"},
// 	// Optional
// 	{"contributor", "coverage", "creator", "date", "description", "format", "publisher", "relation", "rights", "source", "subject", "type"},
// }

// Store all the data in map and then based on the key decide
// if they need tag like <dc:_key_>_value_</dc:_key_> or <meta property="_key_"> _value_ </meta>
func GenerateMetadataSection(writeBuffer *bufio.Writer, items map[string]string) (er error) {
	if _, er = writeBuffer.WriteString("<metadata xmlns:opf=\"http://www.idpf.org/2007/opf\" xmlns:dc=\"http://purl.org/dc/elements/1.1/\">\n"); er != nil {
		return
	}
	for element, value := range items {
		var data string
		switch element {
		case "identifier", "language", "title":
			data = "\t<dc:" + element + ">" + value + "</dc:" + element + ">\n"
		case "contributor", "coverage", "creator", "date", "description", "format", "publisher", "relation", "rights", "source", "subject", "type":
			//TODO: this is incomplete and I have to figure out if mapping from html page's meta section can be mapped directly
		default:
		}
		if _, er = writeBuffer.WriteString(data); er != nil {
			return
		}
	}
	if _, er = writeBuffer.WriteString("</metadata>\n"); er != nil {
		return
	}
	writeBuffer.Flush()
	return
}

// Please make sure to end the section with closing tag
func GeneratePackageStart(writeBuffer *bufio.Writer, uniqueId string) (closing_tag []byte, er error) {
	closing_tag = []byte("</package>")
	if _, er = writeBuffer.WriteString("<package xmlns=\"http://www.idpf.org/2007/opf\" version=\"3.0\""); er != nil {
		return nil, er
	}
	uid := " unique-identifier=\"" + uniqueId + "\">\n"
	if _, er = writeBuffer.WriteString(uid); er != nil {
		return nil, er
	}
	writeBuffer.Flush()
	return closing_tag, nil
}

// ------------------

// Manifest Section of content.opf
type ManifestItem struct {
	FileId    string   `xml:"id,attr"`
	FilePath  string   `xml:"href,attr"`
	MediaType string   `xml:"media-type,attr"`
	XMLName   struct{} `xml:"item"`
}

// TODO: TOC files needs a special it
func GenerateManifestSection(writeBuffer *bufio.Writer, items []ManifestItem) error {
	return generateSection(writeBuffer, "manifest", 1, items)
}

// Spine Section of content.opf
type SpineItem struct {
	Idref   string   `xml:"idref,attr"`
	XMLName struct{} `xml:"itemref"`
}

func GenerateSpineSection(writeBuffer *bufio.Writer, items []SpineItem) error {
	return generateSection(writeBuffer, "spine", 1, items)
}

const XhtmlMime string = "text/html; charset=utf-8"

// type ContentOpfMetaData struct {
// 	// Both can be equal (as far as I understand)
// 	uniqueId, title string
// 	// We will use uuid as identifer
// 	identifier string
// 	// we need short hands like en, de, ru, jp, kr
// 	language string
// }

// Maybe the pages should be more refined which contains the more info
// func GenerateContentOpf(writeBuffer *bufio.Writer, mandatoryMetadata fetch.PageMetadata, assets process.AllAssets) error {
// 	closing, er := GeneratePackageStart(writeBuffer, mandatoryMetadata.Title)
// 	if er != nil {
// 		return er
// 	}
// 	defer func() {
// 		writeBuffer.Write(closing)
// 		writeBuffer.Flush()
// 	}()
// 	// GenerateMetadataSection(writeBuffer, map[string]string{
// 	// 	// Identifier should be uuid
// 	// 	"identifier": mandatoryMetadata.identifier,
// 	// 	"language":   mandatoryMetadata.language,
// 	// 	// Need to find out the the title of main page
// 	// 	"title": mandatoryMetadata.title,
// 	// })
// 	manifestItems := make([]ManifestItem, 0, len(assets.Asssets)+len(assets.XhtmlPages))
// 	spineItems := make([]SpineItem, 0, len(assets.XhtmlPages))
// 	for url, pagedata := range assets.AllAssetStore {
// 		manifestItems = append(manifestItems, ManifestItem{
// 			FileId:    url.EscapedPath(),
// 			FilePath:  url.EscapedPath(),
// 			MediaType: pagedata.MediaType,
// 		})
// 		if pagedata.MediaType == XhtmlMime {
// 			spineItems = append(spineItems, SpineItem{
// 				Idref:   url.EscapedPath(),
// 				XMLName: struct{}{},
// 			})
// 		}
// 	}
// 	GenerateManifestSection(writeBuffer, manifestItems)
// 	// GenerateSpineSection(writeBuffer, spineItems)
// 	return nil
// }

// General Functions

// Need to be very very careful when using this
func removeClose(buf []byte) []byte {
	var idx int = len(buf) - 1
	for buf[idx] != byte('/') {
		idx -= 1
	}
	buf[idx-1] = byte('>')
	buf[idx-2] = byte('/')
	return buf[0:idx]
}

func generateSection[S ManifestItem | SpineItem](writeBuffer *bufio.Writer, sectionName string, indent int, items []S) error {
	l0 := strings.Repeat("\t", indent+0)
	l1 := strings.Repeat("\t", indent+1)

	start := []byte("<" + sectionName + ">\n")
	end := []byte("</" + sectionName + ">\n")

	if _, err := writeBuffer.WriteString(l0); err != nil {
		return err
	}
	if _, err := writeBuffer.Write(start); err != nil {
		return err
	}

	for _, it := range items {
		bytes, er := xml.Marshal(it)

		if er != nil {
			return er
		}

		bytes = removeClose(bytes)

		if _, err := writeBuffer.WriteString(l1); err != nil {
			return err
		}
		if _, err := writeBuffer.Write(bytes); err != nil {
			return err
		}
		if _, err := writeBuffer.WriteRune('\n'); err != nil {
			return err
		}
		writeBuffer.Flush()

	}

	if _, err := writeBuffer.WriteString(l0); err != nil {
		return err
	}
	if _, err := writeBuffer.Write(end); err != nil {
		return err
	}
	writeBuffer.Flush()
	return nil
}
