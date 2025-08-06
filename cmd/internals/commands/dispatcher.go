package commands

import (
	"fmt"
	"io"
	"os"
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
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to get user home directory %w", err)
		}
		appDir, err := EnsureAppDir(homeDir)
		if err != nil {
			return fmt.Errorf("init command error %w", err)
		}
		_, _ = fmt.Fprintf(d.Stdout, "Initialise app in: %s\n", appDir)
		return nil
	case "help":
		return d.printUsage()
	default:
		return fmt.Errorf("unknown command: %s", args[1])
	}
}

func (d Dispatcher) printUsage() error {
	usage := `Gophers and Dragons Character Creator

Usage:
	gnd init 	Initialise app

Examples:
	gnd init # Creates ~/.gnd directory structure

`
	_, _ = fmt.Fprint(d.Stdout, usage)
	return nil
}
