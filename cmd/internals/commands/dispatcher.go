package commands

import "fmt"

type Dispatcher struct{}

func New() *Dispatcher {
	return &Dispatcher{}
}

func (d Dispatcher) Dispatch(args []string) error {
	if len(args) < 2 {
		return d.printUsage()
	}

	switch args[1] {
	case "init":
		fmt.Println("Initialise app")
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
	fmt.Print(usage)
	return nil
}
