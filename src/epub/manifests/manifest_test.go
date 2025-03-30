package manifests_test

import (
	"bufio"
	"bytes"
	"testing"

	"github.com/cloakwiss/cobweb/epub/manifests"
)

func TestManifestSection(t *testing.T) {
	input := []manifests.ManifestItem{
		{FileId: "bookshelf-css", FilePath: "css/bookshelf.css", MediaType: "text/css"},
		{FileId: "book_local-css", FilePath: "css/book_local.css", MediaType: "text/css"},
		{FileId: "book_local-js", FilePath: "scripts/book_local.js", MediaType: "text/javascript"},
		{FileId: "DejaVuSans", FilePath: "fonts/DejaVuSans.ttf", MediaType: "font/opentype"},
		{FileId: "DejaVuSansBold", FilePath: "fonts/DejaVuSans-Bold.ttf", MediaType: "font/opentype"},
		{FileId: "DejaVuSansOblique", FilePath: "fonts/DejaVuSans-Oblique.ttf", MediaType: "font/opentype"},
		{FileId: "DejaVuSansBoldOblique", FilePath: "fonts/DejaVuSans-BoldOblique.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSansMono", FilePath: "fonts/DroidSansMono.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSans", FilePath: "fonts/DroidSans.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSansBold", FilePath: "fonts/DroidSans-Bold.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSerif", FilePath: "fonts/DroidSerif.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSerif-Italic", FilePath: "fonts/DroidSerif-Italic.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSerifBold", FilePath: "fonts/DroidSerif-Bold.ttf", MediaType: "font/opentype"},
		{FileId: "DroidSerif-BoldItalic", FilePath: "fonts/DroidSerif-BoldItalic.ttf", MediaType: "font/opentype"},
		{FileId: "Mincho", FilePath: "fonts/mplus-1c-regular.ttf", MediaType: "font/opentype"},
		{FileId: "Quivira", FilePath: "fonts/Quivira.ttf", MediaType: "font/opentype"},
		{FileId: "cover-image", FilePath: "images/cover.jpg", MediaType: "image/jpeg"},
		{FileId: "apple-image", FilePath: "images/apple-logo-black.jpg", MediaType: "image/jpeg"},
		{FileId: "h1-underline", FilePath: "images/h1-underline.gif", MediaType: "image/gif"},
		{FileId: "joe-image", FilePath: "images/joe.jpg", MediaType: "image/jpeg"},
		{FileId: "wiggly-image", FilePath: "images/WigglyRoad.jpg", MediaType: "image/jpeg"},
		{FileId: "f_0000", FilePath: "f_0000.html", MediaType: "application/xhtml+xml"},
		{FileId: "f_0001", FilePath: "f_0001.html", MediaType: "application/xhtml+xml"},
		{FileId: "f_0002", FilePath: "f_0002.html", MediaType: "application/xhtml+xml"},
		{FileId: "f_0003", FilePath: "f_0003.html", MediaType: "application/xhtml+xml"},
		{FileId: "f_0004", FilePath: "f_0004.html", MediaType: "application/xhtml+xml"},
	}
	buf := make([]byte, 1024)
	out := bytes.NewBuffer(buf)
	writ := bufio.NewWriter(out)
	manifests.GenerateManifestSection(writ, input)
	writ.Flush()
	println(string(out.Bytes()))
}

func TestSpineSection(t *testing.T) {
	buf := make([]byte, 1024)
	out := bytes.NewBuffer(buf)
	writ := bufio.NewWriter(out)
	items := []manifests.SpineItem{{Idref: "cover"}, {Idref: "f_0000"}, {Idref: "f_0001"}, {Idref: "f_0002"}, {Idref: "f_0003"}, {Idref: "f_0004"}, {Idref: "f_0005"}, {Idref: "f_0006"}, {Idref: "f_0007"}, {Idref: "f_0008"}, {Idref: "f_0009"}, {Idref: "f_0010"}, {Idref: "f_0011"}, {Idref: "f_0012"}, {Idref: "f_0013"}, {Idref: "f_0014"}, {Idref: "f_0015"}, {Idref: "f_0016"}, {Idref: "f_0017"}, {Idref: "f_0018"}, {Idref: "f_0019"}, {Idref: "f_0020"}, {Idref: "f_0021"}, {Idref: "f_0022"}, {Idref: "f_0023"}, {Idref: "f_0024"}, {Idref: "f_0025"}, {Idref: "f_0026"}, {Idref: "f_0027"}, {Idref: "f_0028"}, {Idref: "f_0029"}, {Idref: "f_0030"}, {Idref: "f_0031"}, {Idref: "f_0032"}, {Idref: "f_0033"}, {Idref: "f_0034"}, {Idref: "f_0035"}, {Idref: "f_0036"}, {Idref: "f_0037"}, {Idref: "f_0038"}, {Idref: "f_0039"}, {Idref: "f_0040"}, {Idref: "f_0041"}, {Idref: "f_0042"}, {Idref: "f_0043"}, {Idref: "f_0044"}, {Idref: "f_0045"}, {Idref: "f_0046"}, {Idref: "f_0047"}, {Idref: "f_0048"}, {Idref: "f_0049"}, {Idref: "f_0050"}, {Idref: "f_0051"}, {Idref: "f_0052"}, {Idref: "f_0053"}, {Idref: "f_0054"}, {Idref: "f_0055"}, {Idref: "f_0056"}, {Idref: "f_0057"}, {Idref: "f_0058"}, {Idref: "f_0059"}, {Idref: "f_0060"}, {Idref: "f_0061"}, {Idref: "f_0062"}, {Idref: "f_0063"}}
	manifests.GenerateSpineSection(writ, items)
	writ.Flush()
	println(string(out.Bytes()))
}

func TestNewContainer(t *testing.T) {
	buf := make([]byte, 1024)
	out := bytes.NewBuffer(buf)
	writ := bufio.NewWriter(out)
	manifests.NewContainer(writ, "content.opf")
	writ.Flush()
	println(string(out.Bytes()))
}

func TestMetadata(t *testing.T) {
	buf := make([]byte, 1024)
	out := bytes.NewBuffer(buf)
	writ := bufio.NewWriter(out)
	items := map[string]string{"identifier": "Hello", "language": "en", "title": "Goodbye"}
	manifests.GenerateMetadataSection(writ, items)
	writ.Flush()
	println(string(out.Bytes()))
}
