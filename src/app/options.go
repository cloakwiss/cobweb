package app

import (
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
	audio, css, iframe, fonts, js, images, video, metadata bool
	allowDomains, blockDomains                             []url.URL
	output, cookie                                         string
	depth                                                  uint8
	mode                                                   Mode
	timeout                                                time.Duration
}
