## Main logic flow 
+ Create graph structure
  To overlay over the scrappers output and rest of the application
  + Maybe use a scrapper (instead of own library). Some candidates:
    + [colly](https://github.com/gocolly/colly)
    + [geziyor](https://github.com/geziyor/geziyor)
+ Parse out links in page and rewrite it. Some candidate:
  + [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)
  + [goquery](github.com/PuerkitoBio/goquery)
+ Convert to XHTML
  Suggestion is same as above 
  + [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)
    Claude Suggested the following code: (if it works then 🥳🥳)
    ```go
    package main

    import (
        "bytes"
        "fmt"
        "golang.org/x/net/html"
        "strings"
    )

    func main() {
        input := `<div><img src="test.jpg" alt="test"><br><input type="text"></div>`
    
        // Parse HTML
        doc, err := html.Parse(strings.NewReader(input))
        if err != nil {
            panic(err)
        }

        // Function to ensure XHTML compliance
        var ensureXHTML func(*html.Node)
        ensureXHTML = func(n *html.Node) {
            if n.Type == html.ElementNode {
                // Add missing closing tags for void elements
                switch n.Data {
                case "img", "br", "hr", "input", "meta", "link":
                    n.Data = strings.ToLower(n.Data)
                }
            
                // Ensure all attributes are lowercase
                for i := range n.Attr {
                    n.Attr[i].Key = strings.ToLower(n.Attr[i].Key)
                }
            }
        
            // Process children
            for c := n.FirstChild; c != nil; c = c.NextSibling {
                ensureXHTML(c)
            }
        }

        ensureXHTML(doc)

        // Render as XHTML
        var buf bytes.Buffer
        html.Render(&buf, doc)
        fmt.Println(buf.String())
    }     
    ```
  + [goquery](github.com/PuerkitoBio/goquery)

+ Create XML Manifests
  + Manifest
  + Table of content
+ Lay out all the files 

## Good to have
+ Argument Parsing
  [github.com/spf13/cobra](https://github.com/spf13/cobra) for argument parsing
+ Progress Bars
  [bubbletea](https://github.com/charmbracelet/bubbletea) for interfaces 
+ ~C FFI~ I really wish this should not be required


