package zip

import (
	"archive/zip"
	"bytes"
	"log"
	"os"
	"strings"

	"github.com/cloakwiss/cobweb/tidy"
)

// The order of pages is important, so make it array instead of map later
// TODO: add proper date of creation to every file
func WriteTozip(pages map[string][]byte, outputZipFile string) {
	// Create a buffer to write our archive to.
	buf := new(bytes.Buffer)

	// Create a new zip archive.
	w := zip.NewWriter(buf)

	// Add some files to the archive.
	for path, file := range pages {
		if ishtml, newName := cleanHTML(path); ishtml {
			f, err := w.Create(newName)
			if err != nil {
				log.Fatal(err)
			}
			xhtml := tidy.TidyHTML(file)
			_, err = f.Write(xhtml)
			if err != nil {
				log.Fatal(err)
			}
		} else {
			f, err := w.Create(path)
			if err != nil {
				log.Fatal(err)
			}
			_, err = f.Write(file)
			if err != nil {
				log.Fatal(err)
			}
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

// This function makes decision to run the tidy html function or not
// and also suggests new name
func cleanHTML(path string) (bool, string) {
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
