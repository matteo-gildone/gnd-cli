package commands

import (
	"fmt"
	"strings"
)

type HandlerFunc func(d *Dispatcher, args []string) error
type Command struct {
	Name        string
	Description string
	Usage       string
	Examples    []string
	Handler     HandlerFunc
}

func (c Command) String() string {
	var builder strings.Builder
	builder.WriteString(fmt.Sprint("Name:\n"))
	builder.WriteString(fmt.Sprintf("\t%s - %s \n\n", c.Name, c.Description))

	if c.Usage != "" {
		builder.WriteString(fmt.Sprint("Usage:\n"))
		builder.WriteString(fmt.Sprintf("\t%s\n\n", c.Usage))
	}

	if len(c.Examples) > 0 {
		builder.WriteString("Examples:\n")
		for _, example := range c.Examples {
			builder.WriteString(fmt.Sprintf("\t%s\n", example))
		}
	}

	return builder.String()
}
