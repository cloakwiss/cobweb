package zip

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

type Pair struct {
	File  string
	Bytes []byte
}

// The order of pages is important, so make it array instead of map later
// TODO: add proper date of creation to every file
func WriteTozip(pages []Pair, outputZipFile string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Adding mimetype file which should be the first in any epub

	/*{
		file := zip.FileHeader{
			Name:               "mimetype",
			Comment:            "",
			NonUTF8:            true,
			Method:             zip.Store,
			CRC32:              749429103,
			CompressedSize:     20,
			UncompressedSize:   20,
			CompressedSize64:   20,
			UncompressedSize64: 20,
		}
		f, err := w.CreateRaw(&file)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write([]byte("application/epub+zip"))
		if err != nil {
			log.Fatal(err)
		}
	}*/

	// Add some files to the archive.
	for _, pair := range pages {
		f, err := w.Create(pair.File)
		if err != nil {
			log.Fatal(err)
		}
		trimmed := bytes.Trim(pair.Bytes, "0x00")
		_, err = f.Write(trimmed)
		if err != nil {
			log.Fatal(err)
		}
	}

	// Make sure to check the error on Close.
	if err := w.Close(); err != nil {
		log.Fatal(err)
	}
	zipfile, err := os.Create(outputZipFile)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := zipfile.Write(buf.Bytes()); err != nil {
		log.Fatal(err)
	}
	if err := zipfile.Close(); err != nil {
		log.Fatal(err)
	}
}
