package cobweb

import (
	"bufio"
	"encoding/json"
	"encoding/xml"
	"io"
	"os"
	"strings"
)

type ExtractedAttr struct {
	ParentTag string   "json:parent_tag"
	Attribute string   "json:attr"
	Value     []string "json:value"
}

func Start(name string) {
	file := open_file(name)
	if file == nil {
		println("Cannot open file")
		return
	}
	ret := xml.NewDecoder(file)
	attrs := make([]ExtractedAttr, 0)
	for token, err := ret.Token(); err == nil && token != nil; token, err = ret.Token() {
		switch t := token.(type) {
		case xml.StartElement:
			for _, v := range t.Attr {
				attrs = append(attrs, ExtractedAttr{
					ParentTag: t.Name.Local,
					Attribute: v.Name.Local,
					Value:     strings.Split(v.Value, " "),
				})
			}
		}
	}
	out := bufio.NewWriter(os.Stdout)
	encoder := json.NewEncoder(out)
	encoder.Encode(attrs)
	out.Flush()
}

func open_file(name string) io.Reader {
	fd, err := os.Open(name)
	if err != nil {
		println(err)
		return nil
	}
	return bufio.NewReader(fd)
}
