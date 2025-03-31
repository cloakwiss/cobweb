package main

import (
	"log"
	"os"

	"github.com/cloakwiss/cobweb/app"
	"github.com/cloakwiss/cobweb/app/core"
	"github.com/cloakwiss/cobweb/web_ui"
)

func main() {
	// Fine for now, I think so
	if len(os.Args) > 1 {
		args := app.Args()
		output_name := core.Launch(args)
		log.Panicln("The output file: %s\n", output_name)
	} else {
		web_ui.Launch()
	}
}
