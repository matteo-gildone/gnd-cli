package commands

import (
	"fmt"
	"io"
)

type Dispatcher struct {
	Stdout io.Writer
	Stderr io.Writer
}

func New(stdout, stderr io.Writer) *Dispatcher {
	return &Dispatcher{
		Stdout: stdout,
		Stderr: stderr,
	}
}

func (d Dispatcher) Dispatch(args []string) error {
	if len(args) < 2 {
		return d.printUsage()
	}

	switch args[1] {
	case "init":
		_, _ = fmt.Fprintf(d.Stdout, "Initialise app")
		return nil
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func (d Dispatcher) printUsage() error {
	usage := `Gophers and Dragons Character Creator

Usage:
	gnd init 	Initialise app

Examples:
	gnd init

`
	_, _ = fmt.Fprint(d.Stdout, usage)
	return nil
}
