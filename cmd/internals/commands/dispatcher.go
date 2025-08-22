package commands

import (
	"fmt"
	"io"
	"os"
)

type Dispatcher struct {
	Stdout  io.Writer
	Stderr  io.Writer
	HomeDir string
}

func New(stdout, stderr io.Writer) (*Dispatcher, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("failed to get user home directory %w", err)
	}
	d := &Dispatcher{
		Stdout:  stdout,
		Stderr:  stderr,
		HomeDir: homeDir,
	}
	return d, nil
}

func (d *Dispatcher) Dispatch(args []string) error {
	if len(args) < 2 {

		return d.printUsage()
	}
	command := args[1]
	if command == "help" {
		if len(args) >= 3 {
			return d.printCommandHelp(args[2])
		}
		return d.printUsage()
	}

	if cmd, exists := registry.commands[command]; exists {
		return cmd.Handler(d, args)
	}

	return fmt.Errorf("unknown command: %s", command)
}

func (d *Dispatcher) printUsage() error {
	_, _ = fmt.Fprint(d.Stdout, "Gophers and Dragons Character Creator\n\n")
	_, _ = fmt.Fprint(d.Stdout, "Usage:\n")
	_, _ = fmt.Fprint(d.Stdout, "    gnd <command> [args...]\n")
	_, _ = fmt.Fprint(d.Stdout, "    gnd help [command]\n")

	if len(registry.commands) > 0 {
		_, _ = fmt.Fprint(d.Stdout, "Available commands:\n")
		for _, cmd := range registry.GetCommand() {
			_, _ = fmt.Fprintf(d.Stdout, "    %-12s %s\n", cmd.Name, cmd.Description)
		}
	}
	_, _ = fmt.Fprintf(d.Stdout, "    %-12s %s\n\n", "help", "Show help for command")
	_, _ = fmt.Fprint(d.Stdout, "Use gnd help <command> for more information about a command\n")
	return nil
}

func (d *Dispatcher) printCommandHelp(commandName string) error {
	if cmd, exists := registry.commands[commandName]; exists {
		_, _ = fmt.Fprint(d.Stdout, cmd)
		return nil
	}
	return fmt.Errorf("unknown command: %s", commandName)
}
