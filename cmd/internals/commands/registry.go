package commands

import "sort"

type CommandRegistry struct {
	commands map[string]*Command
}

var registry = &CommandRegistry{
	commands: make(map[string]*Command),
}

func Register(cmd *Command) {
	if cmd == nil {
		return
	}

	registry.commands[cmd.Name] = cmd
}

func (r *CommandRegistry) GetCommand() []*Command {
	commands := make([]*Command, 0, len(r.commands))
	for _, cmd := range r.commands {
		commands = append(commands, cmd)
	}

	sort.Slice(commands, func(i, j int) bool {
		return commands[i].Name < commands[j].Name
	})

	return commands
}
