package main

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/app/ui"
	"github.com/cloakwiss/cobweb/epub/zip"
	"github.com/cloakwiss/cobweb/fetch"
)

func main() {
	args := app.Args()
	// fmt.Printf("%+v", args)
	iChan := make(chan app.ApMsg, 10000)

	if len(args.Targets) == 0 {
		println("Early exit")
		return
	}
	fmt.Printf("Targets: %+v\n", args.Targets)

	go func() {
		pages := fetch.Scrapper(args.Targets[0], args, iChan)
		npages := make(map[string][]byte)
		for key, val := range pages {
			// fmt.Printf("Page: %s\n", key.String())
			// Pay attention to the extra `/` at the start of path
			stripped := strings.TrimLeft(key.Path, "/")
			npages[stripped] = val.Data
		}
		zip.WriteTozip(npages, args.Output+".zip")
	}()

	model := ui.NewModel(iChan)
	if _, err := tea.NewProgram(model).Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}

}
