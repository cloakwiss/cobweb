package zip

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

// The order of pages is important, so make it array instead of map later
// TODO: add proper date of creation to every file
func WriteTozip(pages map[string][]byte, outputZipFile string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Adding mimetype file which should be the first in any epub
	{
		file := zip.FileHeader{
			Name:    "mimetype",
			Comment: "",
			NonUTF8: true,
			Method:  zip.Store,
		}
		f, err := w.CreateRaw(&file)
		if err != nil {
			log.Fatal(err)
		}
		f.Write([]byte("application/epub+zip"))
	}
	// Adding META-INF/container.xml which should be the second in epub
	{
		//TODO:
	}

	// Add some files to the archive.
	for path, file := range pages {
		f, err := w.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		_, err = f.Write(file)
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
