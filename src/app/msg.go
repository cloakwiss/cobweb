package app

import (
	"fmt"
)

type MsgCode uint8

const (
	VisitingPage MsgCode = iota
	OnPage
)

type ApMsg struct {
	Code    MsgCode
	Payload string
}

func (m ApMsg) String() (out string) {
	switch m.Code {
	case VisitingPage:
		out = fmt.Sprintf("Visiting Page: %s", m.Payload)
	case OnPage:
		out = fmt.Sprintf("On Page: %s", m.Payload)
	}
	return
}
