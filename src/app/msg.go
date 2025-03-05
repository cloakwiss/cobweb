package app

import (
	"fmt"
)

type MsgCode uint8

const (
	VisitingPage MsgCode = iota
	OnPage
	OnScraped
)

type ApMsg struct {
	Code MsgCode
	URL  string
}

func (m ApMsg) String() (out string) {
	switch m.Code {
	case VisitingPage:
		out = fmt.Sprintf("Visiting Page: %s", m.URL)
	case OnPage:
		out = fmt.Sprintf("On Page: %s", m.URL)
	case OnScraped:
		out = fmt.Sprintf("On Scraped: %s", m.URL)

	}
	return
}
