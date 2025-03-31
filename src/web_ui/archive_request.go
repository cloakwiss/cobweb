package web_ui

import (
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/app/core"
)

type WebOptions struct {
	// Need to work out on how to handle metatdata
	NoAudio, NoCss, NoIframe, NoFonts, NoJs, NoImages, NoVideo, NoMetadata bool
	Targets, AllowDomains, BlockDomains                                     string
	Output, Cookie                                                         string
	Depth                                                                  uint8
	Mode                                                                   string
	Timeout                                                                uint64
}

type Result struct {
	DownloadUrl string
	Message     string
}

// Needed for working with horid json object
func WebOptToOpt(w WebOptions) app.Options {
	target_url, err := url.Parse(strings.TrimSpace(w.Targets))
	if err != nil {
		log.Printf("Error parsing target url: %v\n", err)
	}
	log.Printf("target_url : %v\n", target_url)

	allowed_domains := strings.Split(w.AllowDomains, ",")
	allowed_domain_urls := app.AddUrls(allowed_domains)

	blocked_domains := strings.Split(w.BlockDomains, ",")
	blocked_domains_urls := app.AddUrls(blocked_domains)

	timeout := time.Duration(w.Timeout) * time.Millisecond

	return app.Options{
		NoAudio:      w.NoAudio,
		NoCss:        w.NoCss,
		NoFonts:      w.NoFonts,
		NoJs:         w.NoJs,
		NoImages:     w.NoImages,
		NoVideo:      w.NoVideo,
		Targets:      *target_url,
		AllowDomains: allowed_domain_urls,
		BlockDomains: blocked_domains_urls,
		Output:       w.Output,
		Cookie:       w.Cookie,
		Depth:        w.Depth,
		Timeout:      timeout,
	}
}

func RunApp(args app.Options) Result {

	// Launching the actual Application
	output_name := core.Launch(args)

	return Result{
		DownloadUrl: "/" + output_name,
		Message:     "[Success]",
	}
}
