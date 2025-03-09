package manifests

import (
	"encoding/xml"
	"slices"
	"strings"
)

type ManifestItem struct {
	FileId    string   `xml:"id,attr"`
	FilePath  string   `xml:"href,attr"`
	MediaType string   `xml:"media-type,attr"`
	XMLName   struct{} `xml:"item"`
}

type SpineItem struct {
	Idref   string   `xml:"idref,attr"`
	XMLName struct{} `xml:"itemref"`
}

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

func generateSection[S ManifestItem | SpineItem](sectionName string, indent int, items []S) ([]byte, error) {
	backingBytes := make([]byte, 0, len(items)*128)
	start := []byte("<" + sectionName + ">")
	end := []byte("</" + sectionName + ">")

	backingBytes = append(backingBytes, []byte(strings.Repeat(string('\t'), indent))...)
	backingBytes = append(backingBytes, start...)
	backingBytes = append(backingBytes, byte('\n'))

	for _, it := range items {
		bytes, er := xml.Marshal(it)
		if er != nil {
			return nil, er
		}
		bytes = removeClose(bytes)
		backingBytes = append(backingBytes, []byte(strings.Repeat(string('\t'), indent+1))...)
		backingBytes = append(backingBytes, bytes...)
		backingBytes = append(backingBytes, byte('\n'))

	}

	backingBytes = append(backingBytes, []byte(strings.Repeat(string('\t'), indent))...)
	backingBytes = append(backingBytes, end...)
	backingBytes = append(backingBytes, byte('\n'))

	return slices.Clip(backingBytes), nil
}

func GenerateManifestSection(items []ManifestItem) ([]byte, error) {
	return generateSection("manifest", 1, items)
}
func GenerateSpineSection(items []SpineItem) ([]byte, error) {
	return generateSection("spine", 1, items)
}
