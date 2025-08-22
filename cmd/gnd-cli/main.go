package main

import (
	"fmt"
	"io"
	"os"

	"github.com/matteo-gildone/gnd-cli/cmd/internals/commands"
)

const (
	Success = iota
	Error
)

func main() {
	exitCode := run(os.Args, os.Stdout, os.Stderr)
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func run(args []string, stdout, stderr io.Writer) int {
	commandsDispatcher, _ := commands.New(stdout, stderr)

	if err := commandsDispatcher.Dispatch(args); err != nil {
		_, _ = fmt.Fprintf(stderr, "gnd: %v\n", err)
		return Error
	}
	return Success
}
