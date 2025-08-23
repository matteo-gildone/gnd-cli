package commands

import (
	"fmt"
	"os"
	"path/filepath"
)

const (
	appDir        = ".gnd"
	charactersDir = "characters"
)

func init() {
	Register(&Command{
		Name:        "init",
		Description: "Initialise the application directory structure",
		Usage:       "gnd init",
		Examples: []string{
			"gnd init # Creates ~/.gnd directory structure",
		},
		Handler: handleInit,
	})
}

func handleInit(d *Dispatcher, _ []string) error {
	configDir := filepath.Join(d.HomeDir, appDir)
	_, err := os.Stat(configDir)
	if err == nil {
		_, _ = fmt.Fprintf(d.Stdout, "Reinitialized existing app in %s\n\n", configDir)
	}

	dirs := []string{
		configDir,
		filepath.Join(configDir, charactersDir),
	}

	for _, dir := range dirs {
		_, _ = fmt.Fprintf(d.Stdout, "📁 Creating directory: %s\n", dir)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("failed create directory %s: %w", dir, err)
		}
		_, _ = fmt.Fprintf(d.Stdout, "✅ %s created!\n", dir)
	}

	return nil
}
