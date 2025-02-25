package app

import (
	"log"
	"net/url"
	"time"
)

type Mode uint8

// FIX: make this some thing like enum
const (
	Silent Mode = iota
	Default
	Debug
)

type Options struct {
	// exclude the following or not
	Audio, Css, Iframe, Fonts, Js, Images, Video, Metadata bool
	inputURLs, allowDomains, blockDomains                  []url.URL
	Output, Cookie                                         string
	Depth                                                  uint8
	Mode                                                   Mode
	Timeout                                                time.Duration
}

func parseURL(raw string) (*url.URL, error) {
	u, err := url.Parse(raw)
	if err != nil {
		log.Println("Parsing failed")
		return nil, err
	} else {
		return u, nil
	}
}

func (o *Options) AddUrl(rawURL string) error {
	u, err := parseURL(rawURL)
	if err != nil {
		return err
	}
	o.inputURLs = append(o.inputURLs, *u)
	return nil
}

func (o *Options) AddAllowDomain(rawURLs []string) error {
	for _, raw := range rawURLs {
		u, err := parseURL(raw)
		if err != nil {
			return err
		}
		o.inputURLs = append(o.allowDomains, *u)
	}
	return nil
}
func (o *Options) AddBlockDomain(rawURLs []string) error {
	for _, raw := range rawURLs {
		u, err := parseURL(raw)
		if err != nil {
			return err
		}
		o.inputURLs = append(o.blockDomains, *u)
	}
	return nil
}
