## Main logic flow 
+ Create graph structure
  To overlay over the scrappers output and rest of the application
  + Maybe use a scrapper (instead of own library). Some candidates:
    + [colly](https://github.com/gocolly/colly) using this for project
    + ~[geziyor](https://github.com/geziyor/geziyor)~
+ Parse out links in page and rewrite it. Some candidate:
  + [golang.org/x/net/html](https://pkg.go.dev/golang.org/x/net/html)
  + [goquery](github.com/PuerkitoBio/goquery)
+ Convert to XHTML
  + Will ship `libtidy` along with code to convert to xhtml

+ Create XML Manifests
  + Manifest
  + Table of content
+ Lay out all the files 

## Good to have
+ Argument Parsing
  [github.com/spf13/cobra](https://github.com/spf13/cobra) for argument parsing
+ Progress Bars
  [bubbletea](https://github.com/charmbracelet/bubbletea) for interfaces 


