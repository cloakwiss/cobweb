// both epub 2 and 3 have there own different format
// for now we will work on v3 first the body contain
// nested list within the nav element tag with attribute
// epub:type="toc"
package manifests

import (
	"strings"
)

type Dir struct {
	Path    string
	Files   []File
	SubDirs []Dir
}

type File struct {
	Path string
	Name string
}

func directoryTree(root string, dirlisting []string) Dir {
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
		processedFiles []File = make([]File, len(files))
		processedDirs  []Dir  = make([]Dir, 0)
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

	return Dir{
		Path:    root,
		Files:   processedFiles,
		SubDirs: processedDirs,
	}
}

func RootDirectory(fileslist []string) Dir {
	return directoryTree("", fileslist)
}
