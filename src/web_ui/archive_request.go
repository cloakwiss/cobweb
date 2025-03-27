package web_ui

import (
	"strings"
	"time"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/epub/zip"
	"github.com/cloakwiss/cobweb/fetch"
)

type WebOptions struct {
	// Need to work out on how to handle metatdata
	NoAudio, NoCss, NoIframe, NoFonts, NoJs, NoImages, NoVideo, NoMetadata bool
	Targets, AllowDomains, BlockDomains                                    string
	Output, Cookie                                                         string
	Depth                                                                  uint8
	Mode                                                                   string
	Timeout                                                                uint64
}

type Result struct {
	DownloadUrl string
	Message     string
}

func WebOptToOpt(w WebOptions) app.Options {
	targets := strings.Split(w.Targets, ",")
	target_urls := app.AddUrls(targets)

	allowed_domains := strings.Split(w.AllowDomains, ",")
	allowed_domain_urls := app.AddUrls(allowed_domains)

	blocked_domains := strings.Split(w.BlockDomains, ",")
	blocked_domains_urls := app.AddUrls(blocked_domains)

	timeout := time.Duration(w.Timeout) * time.Millisecond

	return app.Options{
		NoAudio:      w.NoAudio,
		NoCss:        w.NoCss,
		NoIframe:     w.NoIframe,
		NoFonts:      w.NoFonts,
		NoJs:         w.NoJs,
		NoImages:     w.NoImages,
		NoVideo:      w.NoVideo,
		NoMetadata:   w.NoMetadata,
		Targets:      target_urls,
		AllowDomains: allowed_domain_urls,
		BlockDomains: blocked_domains_urls,
		Output:       w.Output,
		Cookie:       w.Cookie,
		Depth:        w.Depth,
		Mode:         app.Mode(app.Debug),
		Timeout:      timeout,
	}
}

func RunApp(args app.Options) Result {
	npages := make(map[string][]byte)
	{
		pages := fetch.Scrapper(args.Targets[0], args)
		for key, val := range pages {
			// fmt.Printf("Page: %s\n", key.String())
			// Pay attention to the extra `/` at the start of path
			stripped := strings.TrimLeft(key.Path, "/")
			npages[stripped] = val.Data
		}
	}
	zip.WriteTozip(npages, args.Output+".zip")

	return Result{
		DownloadUrl: "/" + args.Output + ".zip",
		Message: "Pokemon",
	}
}
