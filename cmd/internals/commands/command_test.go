package commands

import (
	"testing"
)

func TestCommand_String(t *testing.T) {
	testCases := []struct {
		name     string
		command  Command
		expected string
	}{
		{
			name: "minimal command with name and description",
			command: Command{
				Name:        "test",
				Description: "A test command",
			},
			expected: "Name:\n\ttest - A test command \n\n",
		},
		{
			name: "command with name, description and usage",
			command: Command{
				Name:        "deploy",
				Description: "Deploy application",
				Usage:       "deploy [environment]",
			},
			expected: "Name:\n\tdeploy - Deploy application \n\nUsage:\n\tdeploy [environment]\n\n",
		},
		{
			name: "command with name, description and example",
			command: Command{
				Name:        "build",
				Description: "Build project",
				Examples:    []string{"build --prod", "build --dev"},
			},
			expected: "Name:\n\tbuild - Build project \n\nExamples:\n\tbuild --prod\n\tbuild --dev\n",
		},
		{
			name: "command complete of all fields",
			command: Command{
				Name:        "serve",
				Description: "Start the server",
				Usage:       "serve [--port PORT]",
				Examples:    []string{"serve", "serve --port 8080"},
			},
			expected: "Name:\n\tserve - Start the server \n\nUsage:\n\tserve [--port PORT]\n\nExamples:\n\tserve\n\tserve --port 8080\n",
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.command.String()

			if result != tt.expected {
				t.Fatalf("Command.string() = %q, want %q", result, tt.expected)
			}
		})
	}
}
