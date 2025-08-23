package commands

import "fmt"

func init() {
	Register(&Command{
		Name:        "character",
		Description: "List all character",
		Usage:       "gnd character",
		Examples: []string{
			"gnd character # List of characters",
		},
		Handler: handleCreate,
	})
}

func handleCreate(d *Dispatcher, args []string) error {
	_, _ = fmt.Fprintf(d.Stdout, "  📁 Created character: %s\n", args[2])
	return nil
}
