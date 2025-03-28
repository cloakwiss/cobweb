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

func NewContainer(writeBuffer *bufio.Writer, path string) ([]byte, error) {
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
	by, er := xml.MarshalIndent(container, "", "  ")
	if er == nil {
		header = append(header, by...)
		return header, nil
	}
	return nil, er
}

// content.opf

// metadata section of content.opf

var DublinCoreElements = [][]string{
	// Required
	{"identifier", "title", "language"},
	// Optional
	{"contributor", "coverage", "creator", "date", "description", "format", "publisher", "relation", "rights", "source", "subject", "type"},
}

// Store all the data in map and then based on the key decide
// if they need tag like <dc:_key_>_value_</dc:_key_> or <meta property="_key_"> _value_ </meta>
func GenerateMetadateSection(writeBuffer *bufio.Writer, items map[string]string) (er error) {
	if _, er = writeBuffer.WriteString("<metadata xmlns:opf=\"http://www.idpf.org/2007/opf\" xmlns:dc=\"http://purl.org/dc/elements/1.1/\">"); er != nil {
		return
	}
	if _, er = writeBuffer.WriteString("</metadata>"); er != nil {
		return
	}
	for element, value := range items {
		var data string
		switch element {
		case "identifier", "language", "title":
			data = "<dc:" + element + ">" + value + "</dc:" + element + ">"
		case "contributor", "coverage", "creator", "date", "description", "format", "publisher", "relation", "rights", "source", "subject", "type":

		default:
		}
		if _, er = writeBuffer.WriteString(data); er != nil {
			return
		}
	}
	return
}

func GeneratePackageStart(writeBuffer *bufio.Writer, uniqueId string) (er error) {
	if _, er = writeBuffer.WriteString("<package xmlns=\"http://www.idpf.org/2007/opf\" version=\"3.0\""); er != nil {
		return
	}
	uid := " unique-identifier=\"" + uniqueId + "\">"
	if _, er = writeBuffer.WriteString(uid); er != nil {
		return
	}
	return nil
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

	}

	if _, err := writeBuffer.WriteString(l0); err != nil {
		return err
	}
	if _, err := writeBuffer.Write(end); err != nil {
		return err
	}

	return nil
}
