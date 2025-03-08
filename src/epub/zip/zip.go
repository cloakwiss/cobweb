package zip

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
)

// The order of pages is important, so make it array instead of map later
func WriteTozip(pages map[string][]byte, outputZipFile string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

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
