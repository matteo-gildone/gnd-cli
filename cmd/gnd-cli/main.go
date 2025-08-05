package main

import (
	"fmt"
	"github.com/matteo-gildone/gnd-cli/cmd/internals/commands"
	"io"
	"os"
)

func main() {
	exitCode := run(os.Args, os.Stdout, os.Stderr)
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func run(args []string, stdout, stderr io.Writer) int {
	commandsDispatcher := commands.New(stdout, stderr)

	if err := commandsDispatcher.Dispatch(args); err != nil {
		_, _ = fmt.Fprintf(stderr, "Error running the app %v\n", err)
		return 1
	}
	return 0
}
