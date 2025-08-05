package main

import (
	"fmt"
	"github.com/matteo-gildone/gnd-cli/cmd/internals/commands"
	"os"
)

func main() {
	commandsDispatcher := commands.New()

	if err := commandsDispatcher.Dispatch(os.Args); err != nil {
		fmt.Fprintf(os.Stderr, "Error running the app %v\n", err)
		os.Exit(1)
	}
}
