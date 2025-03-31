// both epub 2 and 3 have there own different format
// for now we will work on v3 first the body contain
// nested list within the nav element tag with attribute
// epub:type="toc"
package manifests

import (
	"bufio"
	"encoding/xml"
	"log"
	"strings"
)

type Directory struct {
	Path    string
	Files   []File
	SubDirs []Directory
}

type File struct {
	Path string
	Name string
}

// The Input should be filtered and sorted before hand
func directoryTree(root string, dirlisting []string) Directory {
	var (
		subDirs map[string][]string = make(map[string][]string)
		files   []string
	)

	for i := range dirlisting {
		// removing leading slashes
		if newStr, found := strings.CutPrefix(dirlisting[i], "/"); found {
			dirlisting[i] = newStr
		}
		// classifying
		if strings.ContainsAny(dirlisting[i], "/") {
			dir, rest, _ := strings.Cut(dirlisting[i], "/")
			if _, found := subDirs[dir]; !found {
				subDirs[dir] = make([]string, 0)
			}
			subDirs[dir] = append(subDirs[dir], rest)
		} else {
			files = append(files, dirlisting[i])
		}

	}

	var (
		processedFiles []File      = make([]File, len(files))
		processedDirs  []Directory = make([]Directory, 0)
	)

	for dir, sub := range subDirs {
		var localpath string
		if root != "" {
			localpath = root + "/" + dir
		} else {
			localpath = dir
		}
		processedDirs = append(processedDirs, directoryTree(localpath, sub))
	}
	for i := range files {
		processedFiles[i] = File{Name: files[i], Path: root}
	}

	if len(processedDirs) == 0 {
		processedDirs = nil
	}
	if len(processedFiles) == 0 {
		processedFiles = nil
	}

	return Directory{
		Path:    root,
		Files:   processedFiles,
		SubDirs: processedDirs,
	}
}

func GenerateDirectoryTree(fileslist []string) Directory {
	return directoryTree("root", fileslist)
}

func MarshalToc(dir Directory, writeBuffer *bufio.Writer) {
	header := []string{
		"<!DOCTYPE html>",
		`<html xmlns="http://www.w3.org/1999/xhtml" xml:lang="en" lang="en" xmlns:epub="http://www.idpf.org/2007/ops">`,
		"  <body>",
		"    <nav epub:type=\"toc\">",
		"      <h1>Table of Contents</h1>",
		"      <ol>",
	}
	footer := []string{
		"      </ol>",
		"    </nav>",
		"  </body>",
		"</html>",
	}
	for _, l := range header {
		writeBuffer.WriteString(l + "\n")
	}
	writeBuffer.Flush()
	marshalToc(dir, 5, writeBuffer)
	for _, l := range footer {
		writeBuffer.WriteString(l + "\n")
	}
	writeBuffer.Flush()
}

func marshalToc(dir Directory, indent int, writeBuffer *bufio.Writer) {
	l0 := strings.Repeat("  ", indent+0)
	l1 := strings.Repeat("  ", indent+1)
	l2 := strings.Repeat("  ", indent+2)

	writeBuffer.WriteString(l0 + "<li>")
	defer func() {
		writeBuffer.WriteString(l0 + "</li>\n")
		writeBuffer.Flush()
	}()

	writeBuffer.WriteString("<span>")
	if er := xml.EscapeText(writeBuffer, []byte(dir.Path)); er != nil {
		log.Fatal("Marshalling Toc failed during writing Path")
	}
	writeBuffer.WriteString("</span>\n")

	writeBuffer.WriteString(l1 + "<ol>\n")
	writeBuffer.Flush()

	for _, subdir := range dir.SubDirs {
		marshalToc(subdir, indent+1, writeBuffer)
	}

	for _, file := range dir.Files {
		writeBuffer.WriteString(l2 + "<li><a href=\"")
		writeBuffer.WriteString(file.Path + "/" + file.Name)
		writeBuffer.WriteString("\">")
		if er := xml.EscapeText(writeBuffer, []byte(file.Name)); er != nil {
			log.Fatal("Marshalling Toc failed during writing Path")
		}
		writeBuffer.WriteString("</a></li>\n")
	}
	writeBuffer.WriteString(l1 + "</ol>\n")
}
