package ui

// This interface allows to print to TUI
// like log
type Printer interface {
	MsgType() string
	Paylod() string
}
