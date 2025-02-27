package main

import (
	"fmt"
	"os"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/app/ui"
	"github.com/cloakwiss/cobweb/epub/zip"
	"github.com/cloakwiss/cobweb/fetch"
)

// This is rust's documentation which is hosted by caddy to run test
const target = "http://localhost:8080/std/index.html"

func Test_Mainloop(t *testing.T) {
	iChan := make(chan app.ApMsg, 1000)
	// go app.Print(iChan)

	go func() {
		pages := fetch.Mainloop(target, 1, iChan)
		npages := make(map[string][]byte)
		for key, val := range pages {
			// fmt.Printf("Page: %s\n", key.String())
			// Pay attention to the extra `/` at the start of path
			npages[key.Path] = val
		}
		zip.WriteTozip(npages, "output.zip")
	}()

	model := ui.NewModel(iChan)
	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
