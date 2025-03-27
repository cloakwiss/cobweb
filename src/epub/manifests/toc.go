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
		localpath := root + "/" + dir
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
	return directoryTree("", fileslist)
}

func MarshalToc(dir Directory, writeFile *bufio.Writer) {
	marshalToc(dir, 0, writeFile)
}

func marshalToc(dir Directory, indent int, writeFile *bufio.Writer) {
	space := "    "
	l0 := strings.Repeat(space, indent+0)
	l1 := strings.Repeat(space, indent+1)
	l2 := strings.Repeat(space, indent+2)

	writeFile.WriteString(l0 + "<li>")
	defer writeFile.WriteString(l0 + "</li>\n")

	writeFile.WriteString("<span>")
	if er := xml.EscapeText(writeFile, []byte(dir.Path)); er != nil {
		log.Fatal("Marshalling Toc failed during writing Path")
	}
	writeFile.WriteString("</span>\n")

	writeFile.WriteString(l1 + "<ol>\n")
	writeFile.Flush()

	for _, subdir := range dir.SubDirs {
		marshalToc(subdir, indent+1, writeFile)
	}

	for _, file := range dir.Files {
		writeFile.WriteString(l2 + "<li><a href=\"")
		writeFile.WriteString(file.Path + "/" + file.Name)
		writeFile.WriteString("\">")
		if er := xml.EscapeText(writeFile, []byte(file.Name)); er != nil {
			log.Fatal("Marshalling Toc failed during writing Path")
		}
		writeFile.WriteString("</a></li>\n")
	}
	writeFile.WriteString(l1 + "</ol>\n")
	writeFile.Flush()

}
