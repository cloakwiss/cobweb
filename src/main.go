package main

import (
	"fmt"

	"github.com/cloakwiss/cobweb/app"
)

func main() {
	args := app.Args()
	fmt.Printf("%+v", args)
}
