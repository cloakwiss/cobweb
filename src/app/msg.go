package app

import "fmt"

type MsgCode uint8

const (
	VisitingPage MsgCode = iota
	OnPage
)

type Msg struct {
	Code    MsgCode
	Payload string
}

func (m Msg) String() (out string) {
	switch m.Code {
	case VisitingPage:
		out = fmt.Sprintf("Visiting Page: %s", m.Payload)
	case OnPage:
		out = fmt.Sprintf("On Page: %s", m.Payload)
	}
	return
}

func Print(input <-chan Msg) {
	// for {
	// 	inp, open := <-input
	// 	println(inp.String())
	// 	if !open {
	// 		break
	// 	}
	// }
	for msg := range input {
		println(msg.String())
	}
}
